package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_Getpagesize(t *testing.T) {
	pageSize := os.Getpagesize() // 内存页
	fmt.Println(pageSize)
}

func Test_SameFile(t *testing.T) { // 判断是否是同一类型的文件
	finfo1, err := os.Stat("./data/data.txt")
	if err != nil {
		panic(err)
	}
	finfo2, err := os.Stat("./data/001_02_1.wav")
	if err != nil {
		panic(err)
	}

	flag := os.SameFile(finfo1, finfo2)
	fmt.Println(flag)
}

func Test_FileMode(t *testing.T) {
	mp := os.ModePerm
	fmt.Println(mp.String())
	fmt.Println(mp.IsDir())
	fmt.Println(mp.IsRegular())
	fmt.Println(mp.Perm())
}
