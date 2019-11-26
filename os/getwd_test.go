package os

import (
	"fmt"
	"os"
	"syscall"
	"testing"
)

func Test_Getwd(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)

	wd, err := syscall.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(wd)
}
