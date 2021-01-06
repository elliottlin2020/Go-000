package counter

import (
	"sync"
	"time"
)

type Counter interface {
	Add(n int) int
}

type windowSlidingCounter struct {
	start       time.Time
	window      time.Duration
	granularity int64
	generation  int64
	slots       []slot
	total       int
	mu sync.Mutex
}

type slot struct {
	iter int64
	n    int
}

func New(w time.Duration, n int64) Counter {
	if w < time.Millisecond || n < 1 || n > 1000 {
		return nil
	}

	return &windowSlidingCounter{
		start:       time.Now(),
		window:      w,
		granularity: n,
		generation:  0,
		slots:       make([]slot, n),
		mu: sync.Mutex{},
	}
}

func (w *windowSlidingCounter) Add(n int) int {
	w.mu.Lock()
	elapsed := time.Now().Sub(w.start).Microseconds()
	w.generation = elapsed / w.window.Microseconds()
	k := w.window.Microseconds() / w.granularity
	s := (elapsed % w.window.Microseconds()) / k

	if w.generation != w.slots[s].iter {
		w.newGeneration()
	}
	w.slots[s].n += n

	w.total += n
	w.mu.Unlock()

	return w.total
}

func (w *windowSlidingCounter) newGeneration() {
	w.slots = make([]slot, w.granularity)
	for i := range w.slots {
		w.slots[i].iter = w.generation
	}
	w.total = 0
}
