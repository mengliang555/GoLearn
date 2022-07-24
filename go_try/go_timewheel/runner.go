package go_timewheel

import (
	"context"
	"time"
)

type TimeWheelRunner struct {
	currentIndex int
	totalSize    int
	slotList     *slot
	Running      bool
}

type slot struct {
	data []*ListWithRoot
}

func (t *TimeWheelRunner) StartCircle(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func(tr *time.Ticker) {
		for {
			select {
			case <-tr.C:
				doLoop(func(ctx context.Context) {
					t.currentIndex = (t.currentIndex + 1) % t.totalSize
				})
			}
		}
	}(ticker)
}

func doLoop(behave func(ctx context.Context)) {

}
