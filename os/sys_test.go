package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_sys(t *testing.T) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(name) // DESKTOP-J6DGE5E
}
