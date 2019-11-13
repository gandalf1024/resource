package io

import (
	"fmt"
	"io"
	"strconv"
	"sync"
	"testing"
)

func Test_Pipe(t *testing.T) {
	pr, pw := io.Pipe()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() { //读出
		bys := make([]byte, 100)
		n, _ := pr.Read(bys)
		fmt.Println(n)
		fmt.Println(string(bys))
		wg.Done()
	}()

	go func() { //写入
		n, _ := pw.Write([]byte("123456"))
		fmt.Println(n)
		wg.Done()
	}()

	wg.Wait()

}

func Test_Pipe2(t *testing.T) {
	pr, pw := io.Pipe()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() { //读出
		index := 0
		for index < 100000 {
			bys := make([]byte, 100)
			n, _ := pr.Read(bys)
			fmt.Println(n)
			fmt.Println(string(bys))
			index++
		}
		wg.Done()
	}()

	go func() { //写入
		index := 0
		for index < 100000 {
			n, _ := pw.Write([]byte("123456:" + strconv.Itoa(index)))
			fmt.Println(n)
			index++
		}
		wg.Done()
	}()

	wg.Wait()

}

func Test_Pipe_for(t *testing.T) {
	index := 0
	b := make([]byte, 10)
	for once := true; once || len(b) > 0; once = false { //只有  (左||右)  都为false才不进入
		fmt.Println("-----")
		index++
		if index == 3 {
			b = nil
		}
	}

	a := 1 != 1
	v := 2 != 2

	for a || v {
		fmt.Println("11111111")
	}
}

func Test_Copy_Demo(t *testing.T) {
	a := []byte("1111")
	b := []byte("22999922")
	n := copy(a, b)
	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(n)

}
