package io

import (
	"fmt"
	"io"
	"testing"
)

type S struct {
	data string
}

func (str *S) WriteString(s string) (n int, err error) {
	str.data += s
	return len(s) * 10, nil
}

func (str *S) Write(p []byte) (n int, err error) {
	str.data += string(p)
	return len(p) * 5, nil
}

func Test_StringWriter(t *testing.T) {
	str := &S{}
	n, err := io.WriteString(str, "你hi撒旦广泛的事故发生的")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

type R struct {
	bys []byte
}

func (r *R) Read(p []byte) (n int, err error) {
	r.bys = append(r.bys, p...)
	return len(p), nil
}

func Test_ReadAtLeast(t *testing.T) {
	r := &R{}
	r.bys = append(r.bys, 22)
	n, err := io.ReadAtLeast(r, []byte("11212312"), 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	fmt.Println(len(r.bys))

}
