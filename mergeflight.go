package mergeflight

import (
	"sync"
	"time"
)

type call struct {
	Cnt     int
	CntDone chan struct{}
	Args    []interface{}

	Done chan struct{}
	Val  interface{}
	Err  error
}

func newCall(arg interface{}) *call {
	return &call{
		Cnt:     1,
		Args:    []interface{}{arg},
		CntDone: make(chan struct{}),
		Done:    make(chan struct{}),
	}
}

type Merge struct {
	maxCnt  int
	timeWin time.Duration

	c  *call
	lk sync.Mutex
}

// New make Merge object with count window and time window
func New(maxCnt int, timeWin time.Duration) *Merge {
	return &Merge{maxCnt: maxCnt, timeWin: timeWin}
}

// Do pass one arg for batch function and call later, returns the batch result
func (m *Merge) Do(arg interface{}, f func(args []interface{}) (interface{}, error)) (interface{}, error) {
	var c *call
	m.lk.Lock()
	if m.c != nil {
		// has current call
		c = m.c

		// append current arg
		c.Args = append(c.Args, arg)
		c.Cnt++
		reach := c.Cnt >= m.maxCnt
		if reach {
			// reach the count window limit,
			// reset current call to nil
			m.c = nil
		}
		m.lk.Unlock()

		if reach {
			// notify the first call goroutine
			close(c.CntDone)
		}

		<-c.Done
		return c.Val, c.Err
	}
	// make a new call
	c = newCall(arg)
	m.c = c
	m.lk.Unlock()

	to := time.NewTimer(m.timeWin)

	select {
	case <-c.CntDone:
		// reach the count window limit
		to.Stop()
	case <-to.C:
		// reach the time window
		// than reset current call to nil
		m.lk.Lock()
		// check mine corrent
		if m.c == c {
			m.c = nil
		}
		m.lk.Unlock()
	}

	// call function with batch args
	c.Val, c.Err = f(c.Args)
	close(c.Done)
	return c.Val, c.Err
}
