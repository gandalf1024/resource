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

func Test_Lookup(t *testing.T) {
	u, err := user.Lookup("ASUS")
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Name)
	fmt.Println(u.Gid)
	fmt.Println(u.HomeDir)
}

func Test_LookupId(t *testing.T) {
	user, err := user.LookupId("S-1-5-21-4093140098-1533531953-3217901389-1001") // windows (whoami /USER)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.HomeDir)
}

func Test_GroupIds(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	strs, err := u.GroupIds()
	if err != nil {
		panic(err)
	}
	for _, v := range strs {
		fmt.Println(v)
	}
}

func Test_LookupGroup(t *testing.T) {
	g, err := user.LookupGroup("XXXX") // 用户组
	if err != nil {
		panic(err)
	}
	fmt.Println(g.Gid)
	fmt.Println(g.Name)
}
