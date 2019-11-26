package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_Executable(t *testing.T) {
	cutable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println(cutable)
}
