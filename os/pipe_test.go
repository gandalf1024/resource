package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_Pipe(t *testing.T) {
	rf, wf, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			bys := make([]byte, 10)
			_, err = rf.Read(bys)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bys))
		}
	}()

	go func() {
		for {
			_, err := wf.Write([]byte("5235464356"))
			if err != nil {
				panic(err)
			}
		}
	}()

	select {}

}

func Test_Link(t *testing.T) {
	err := os.Link("./data/data.txt", "999") // 类似复制文件
	if err != nil {
		panic(err)
	}
}
