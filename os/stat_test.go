package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_Stat(t *testing.T) { // 获取文件状态
	finfo, err := os.Stat("./data/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(finfo.Name())

}

func Test_Lstat(t *testing.T) {
	finfo, err := os.Lstat("./data")
	if err != nil {
		panic(err)
	}
	fmt.Println(finfo.Name())
}
