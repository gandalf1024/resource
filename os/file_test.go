package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_Open(t *testing.T) {
	f, err := os.Open("./data")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Name())
}
