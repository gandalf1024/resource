package sync

import (
	"fmt"
	"sync"
	"testing"
)

func Test_pool(t *testing.T) {
	p := sync.Pool{}
	p.Put("1")
	result := p.Get()
	fmt.Println(result)
}
