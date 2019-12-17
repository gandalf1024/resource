package sync

import (
	"sync"
	"testing"
)

func Test_Add(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

}
