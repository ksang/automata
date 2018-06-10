package rrul

import (
	"fmt"
	"log"
	"net"
	"time"

	fastping "github.com/ksang/go-fastping"
)

// PingProbe sends ICMP packets and probing round trip latency
func PingProbe(target string, resCh chan []DataPoint, quitCh chan struct{}) {
	res := []DataPoint{}
	p := fastping.NewPinger()
	p.MaxRTT = 500 * time.Millisecond
	ra, err := net.ResolveIPAddr("ip4:icmp", target)
	if err != nil {
		fmt.Println("++ERROR sending ICMP:", err)
		resCh <- res
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		dp := DataPoint{
			Value: float64(rtt.Nanoseconds()) / float64(time.Millisecond),
			Unit:  "ms",
			Time:  float64(time.Now().Unix()),
		}
		res = append(res, dp)
	}
	p.RunLoop()
	select {
	case <-p.Done():
		if err := p.Err(); err != nil {
			log.Fatalf("Ping failed: %v", err)
		}
	case <-quitCh:
		break
	}
	resCh <- res
}
