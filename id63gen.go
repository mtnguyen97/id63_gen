package id63gen

import (
	"math/rand"
	"sync"
	"time"
)

const (
	bits  = 15
	max   = 1 << bits
	bmask = max - 1
	tmask = 0x7fffffffffffffff ^ bmask
)

var (
	previousTicks int64
	pos           int
	box           [max]int16
	mut           sync.Mutex
	resetOnce     sync.Once
)

func reset() {
	previousTicks = 0
	pos = 0
	for i := 0; i < max; i++ {
		box[i] = int16(i)
	}
}

// Next returns an ID
func Next() int64 {
	resetOnce.Do(reset)
	mut.Lock()
	defer mut.Unlock()
	now := time.Now().UnixNano()
	ticks := now & tmask
	if ticks != previousTicks {
		previousTicks = ticks
		pos = 0
		return next()
	}
	if pos < max {
		return next()
	}
	time.Sleep(time.Duration(max - now&bmask))
	previousTicks = ticks + max
	pos = 0
	return next()
}

func next() (id int64) {
	swapAt := pos + rand.Intn(max-pos)
	box[pos], box[swapAt] = box[swapAt], box[pos]
	id = previousTicks | int64(box[pos])
	pos++
	return
}
