package sync

import (
	"fmt"
	"sync"
	"testing"
)

func Test_do(t *testing.T) {
	var m sync.Once
	m.Do(func() {
		fmt.Println("hello")
	})

}
