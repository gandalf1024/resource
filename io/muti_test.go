package io

import (
	"fmt"
	"io"
	"testing"
)

func Test_MultiReader(t *testing.T) {
	r1 := &R{}
	r2 := &R{}

	r1.bys = []byte("123456-")
	r2.bys = []byte("963258")

	reader := io.MultiReader(r1, r2)
	bys := make([]byte, 100)
	n, err := reader.Read(bys)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(bys))
}

func Test_Write(t *testing.T) {
	w1 := &W{}
	w2 := &W{}

	w1.bys = []byte("123456-")
	w2.bys = []byte("963258-")

	writer := io.MultiWriter(w1, w2)
	n, err := writer.Write([]byte("852456"))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(w1.bys))
	fmt.Println(string(w2.bys))
}
