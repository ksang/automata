package rrul

import (
	"testing"
	"time"
)

func TestPingProbe(t *testing.T) {
	quitCh := make(chan struct{}, 1)
	resCh := make(chan []DataPoint, 1)
	go PingProbe("127.0.0.1", resCh, quitCh)
	go func(qc chan struct{}) {
		time.Sleep(10 * time.Second)
		qc <- struct{}{}
	}(quitCh)
	res := <-resCh
	for _, r := range res {
		t.Logf("PingProbe result: %#v", r)
	}
}
