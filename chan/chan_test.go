package _chan

import (
	"fmt"
	"testing"
)

func Test_chan(t *testing.T) {
	ch := make(chan int, 1)
	close(ch)
	rev := <-ch // 从关闭的通道取数据是空值
	fmt.Println(rev)

	ch <- 1 // 发送数据给关闭的通道会引发panic
}

func Test_chan2(t *testing.T) {
	ch := make(chan int, 1)
	//close(ch)

	go func() {
		i := 0
		for i < 100 {
			ch <- i
			i++
		}
		//close(ch)
	}()

	for v := range ch { // 永久阻塞  因为没有close 27 行代码
		fmt.Println(v)
	}

}
