package os

import (
	"fmt"
	"os/user"
	"testing"
)

func Test_Current(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Name)
	fmt.Println(u.Gid)
	fmt.Println(u.HomeDir)
	fmt.Println(u.Username)
}
