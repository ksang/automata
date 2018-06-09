package rrul

import (
	"context"
	"fmt"
	"io/ioutil"
	"os/exec"
)

var (
	netperfCmd    = "netperf"
	tcpUpload     = "TCP_STREAM"
	tcpDownload   = "TCP_MAERTS"
	udpRR         = "UDP_RR"
	netperfParams = " -P 0 -v 0 -D -0.20 -4 -Y CS1,CS1 -H %s -p %d -t %s " +
		"-l %d -F /dev/urandom -f m " +
		" -- -k THROUGHPUT,LOCAL_CONG_CONTROL,REMOTE_CONG_CONTROL,TRANSPORT_MSS," +
		"LOCAL_TRANSPORT_RETRANS,REMOTE_TRANSPORT_RETRANS,LOCAL_SOCKET_TOS,REMOTE_SOCKET_TOS," +
		"DIRECTION,ELAPSED_TIME,PROTOCOL,LOCAL_SEND_SIZE,LOCAL_RECV_SIZE,REMOTE_SEND_SIZE," +
		"REMOTE_RECV_SIZE -P %d"
)

type netperf struct {
	cfg    Config
	data   Result
	quitCh chan struct{}
}

func (n *netperf) start() error {
	if _, err := exec.LookPath(netperfCmd); err != nil {
		return err
	}
	ctx := context.Background()
	return n.runTests(ctx)
}

func (n *netperf) runTests(ctx context.Context) error {
	tcpUploadCmd := netperfCmd + fmt.Sprintf(netperfParams,
		n.cfg.Host, n.cfg.ControlPort, tcpUpload, n.cfg.Seconds, n.cfg.DataPort)
	tcpDownloadCmd := netperfCmd + fmt.Sprintf(netperfParams,
		n.cfg.Host, n.cfg.ControlPort, tcpDownload, n.cfg.Seconds, n.cfg.DataPort)
	udpRRCmd := netperfCmd + fmt.Sprintf(netperfParams,
		n.cfg.Host, n.cfg.ControlPort, udpRR, n.cfg.Seconds, 0)
	tcpuCh := make(chan []byte, 1)
	tcpdCh := make(chan []byte, 1)
	udpCh := make(chan []byte, 1)

	go dispatchJob(ctx, tcpUploadCmd, tcpuCh)
	go dispatchJob(ctx, tcpDownloadCmd, tcpdCh)
	go dispatchJob(ctx, udpRRCmd, udpCh)
	tcpu := <-tcpuCh
	tcpd := <-tcpdCh
	udp := <-udpCh
	tcpuRes, tcpuT, _ := MarshalOutput(tcpu)
	tcpdRes, tcpdT, _ := MarshalOutput(tcpd)
	udpRes, udpT, _ := MarshalOutput(udp)
	n.data.TCPDownload = tcpuRes
	n.data.TCPUpload = tcpdRes
	n.data.UDPRR = udpRes
	n.data.TCPDownloadThroughput = tcpdT
	n.data.TCPUploadThroughput = tcpuT
	n.data.UDPRRThroughput = udpT
	return nil
}

func dispatchJob(ctx context.Context, cmd string, out chan []byte) {
	ret, _ := runCmd(ctx, cmd)
	out <- ret
}

func runCmd(pctx context.Context, cmd string) ([]byte, error) {
	ctx, cancel := context.WithCancel(pctx)
	defer cancel()
	fmt.Printf("++Running command:\n  %s\n", cmd)
	c := exec.CommandContext(ctx, "bash", "-c", cmd)
	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()
	if err := c.Start(); err != nil {
		return nil, err
	}
	out, _ := ioutil.ReadAll(stdout)
	slurp, _ := ioutil.ReadAll(stderr)
	if err := c.Wait(); err != nil {
		fmt.Printf("Stderr:\n\t%s\n", slurp)
		return nil, err
	}
	return out, nil
}
