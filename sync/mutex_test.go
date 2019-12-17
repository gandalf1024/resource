package sync

import (
	"fmt"
	"sync"
	"testing"
)

func Test_mutex(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	fmt.Println("---")
	mu.Unlock()
}

func Test_itoa(t *testing.T) {
	fmt.Println(1 << 0)
}

type User struct {
	sync.Mutex
	name string
}

func Test_dead(t *testing.T) {
	u1 := &User{name: "test"}
	u1.Lock()
	defer u1.Unlock()

	tmp := *u1
	u2 := &tmp
	u2.Mutex = sync.Mutex{} // 没有这一行就会死锁

	fmt.Printf("%#p\n", u1)
	fmt.Printf("%#p\n", u2)

	u2.Lock()
	defer u2.Unlock()
}

func Test_lock(t *testing.T) {
	var n sync.Mutex
	n.Lock()
}
