package ioutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Test_Reader(t *testing.T) {
	r := &R{}
	r.bys = []byte("13452364")

	bys, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bys))
}

func Test_Sli(t *testing.T) {
	bys := []byte("123456789")
	Change(bys[2:])
	fmt.Println(string(bys))
	Change2(bys)
	fmt.Println(string(bys))

}

func Change(bb []byte) {
	copy(bb[2:], "9687461561")
}

func Change2(bb []byte) {
	copy(bb[2:], "9687461561")
}

func Test_KK(t *testing.T) {
	fmt.Println(^uint(0) >> 1) // int最大值
	fmt.Println(^uint(0) >> 0) // uint最大值
	fmt.Println(^uint(0))      // uint最大值
}

var 你好 = "dsddd"

func Test_ReadFile(t *testing.T) {
	fileBys, err := ioutil.ReadFile("./data/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileBys))
}

func Test_WriteFile(t *testing.T) {
	err := ioutil.WriteFile("./data/xx.txt", []byte("123456999889989"), os.ModeAppend)
	if err != nil {
		panic(err)
	}
}

func Test_ReadDir(t *testing.T) {
	finfo, err := ioutil.ReadDir("./data")
	if err != nil {
		panic(err)
	}
	for _, v := range finfo {
		fmt.Println(v.Size())
		fmt.Println(v.IsDir())
		fmt.Println()
	}
}
