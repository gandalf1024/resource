package os

import (
	"os"
	"testing"
)

func Test_MkdirAll(t *testing.T) {
	err := os.MkdirAll("./data3/2/3/4", os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func Test_RemoveAll(t *testing.T) {
	err := os.RemoveAll("./data3")
	if err != nil {
		panic(err)
	}
}

func Test_IsPathSeparator(t *testing.T) {
	println(os.IsPathSeparator('\\'))
}
