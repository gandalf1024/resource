package context

import (
	"fmt"
	"testing"
	"time"
)

type emptyTime int64

func (*emptyTime) Deadline() (deadline time.Time, ok bool) {
	return
}

func Test_emptyTime(t *testing.T) {
	var em emptyTime
	resTime, status := em.Deadline()
	fmt.Println(resTime, status)

	v := time.Time{}
	fmt.Println(v)
	fmt.Println("--->>", v.Year())
}

func Test_befor(t *testing.T) {
	now := time.Now()
	empty := time.Time{}
	println(empty.Before(now))
}
