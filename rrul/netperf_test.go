package rrul

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		cmd  string
		args string
	}{
		{
			"ifconfig",
			"",
		},
		{
			"ls -a -l",
			"",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for _, te := range tests {
		ret, err := runCmd(ctx, te.cmd)
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
		t.Logf("result: %s", string(ret))
	}
}

func TestNetPerfCmd(t *testing.T) {
	tests := []struct {
		cmd  string
		args string
	}{
		{
			netperfCmd + fmt.Sprintf(netperfParams, "127.0.0.1", 12865, tcpUpload, 3, 12866),
			"",
		},
		{
			netperfCmd + fmt.Sprintf(netperfParams, "127.0.0.1", 12865, tcpDownload, 3, 12866),
			"",
		},
		{
			netperfCmd + fmt.Sprintf(netperfParams, "127.0.0.1", 12865, udpRR, 3, 0),
			"",
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		runCmd(ctx, "netserver -D")
	}()
	time.Sleep(time.Second)
	for _, te := range tests {
		ret, err := runCmd(ctx, te.cmd)
		if err != nil {
			t.Errorf("error: %v\n", err)
		}
		fmt.Println(string(ret))
	}
}
