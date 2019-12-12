package time

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Test_Now(t *testing.T) {
	ti := time.Now()
	fmt.Println(ti)

	utc := ti.UTC()
	fmt.Println(utc) // 格林日志时间
}

func Test_After(t *testing.T) {
	hou := time.Now()
	qian := time.Now().Add(-time.Hour * 2)
	flag := hou.Before(qian)
	fmt.Println(flag)
	flag = hou.After(qian)
	fmt.Println(flag)

	fmt.Println(hou)
	fmt.Println(qian)
}

func Test_Equal(t *testing.T) {
	now1 := time.Now()
	now2 := time.Now()
	fmt.Println(now1.Equal(now2))
	now1 = now1.Add(-time.Hour * 2)
	fmt.Println(now1.Equal(now2))
}

func Test_Month(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Month().String())
}

func Test_IsZero(t *testing.T) {
	now := time.Now()
	fmt.Println(now.IsZero())
}

func Test_hao(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(12)

	for i := 0; i < 6; i++ {
		go func() {
			fmt.Println("T1:", i)
			wg.Done()
		}()
	}

	for i := 0; i < 6; i++ {
		go func(i int) {
			fmt.Println("T2:", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
