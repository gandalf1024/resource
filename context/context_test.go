package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_WithCancel(t *testing.T) {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done(): // close后进入
					close(dst)
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 可以多次close

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 50 {
			cancel() // close chan操作
		}
	}

	//gn := gen(ctx)
	//for {
	//	select {
	//	case n,ok := <-gn:
	//		if !ok {
	//			return
	//		}
	//		fmt.Println(n)
	//		if n == 50 {
	//			cancel()
	//		}
	//	}
	//}
}

func Test_WithDeadline(t *testing.T) {
	d := time.Now().Add(1 * time.Second)
	d2 := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d2)
	ctx2, cancel2 := context.WithDeadline(ctx, d)
	defer cancel()
	defer cancel2()

	index := int64(0)
	for {
		if index >= 3 {
			break
		}
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println("ctx...1s ")
		case <-ctx2.Done():
			fmt.Println("ctx2...2s ")
		}
		index++
	}
}

func Test_WithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func Test_WithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))
}
