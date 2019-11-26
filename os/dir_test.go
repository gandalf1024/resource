package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_Readdir(t *testing.T) {
	f, err := os.Open("./data")
	if err != nil {
		panic(err)
	}

	infos, err := f.Readdir(9)
	if err != nil {
		panic(err)
	}

	for _, v := range infos {
		fmt.Println(v.Name())
	}

}
