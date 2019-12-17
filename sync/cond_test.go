package sync

import (
	"sync"
	"testing"
)

func Test_NewCond(t *testing.T) {
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)
	cond.L.Lock()
	cond.Broadcast()
	cond.Signal()
	cond.L.Unlock()
	//cond.Wait()
}
