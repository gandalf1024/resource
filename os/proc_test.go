package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_Getuid(t *testing.T) {
	uid := os.Geteuid()
	fmt.Println(uid)
}

func Test_Geteuid(t *testing.T) {
	euid := os.Geteuid()
	fmt.Println(euid)
}

func Test_Getgid(t *testing.T) {
	gid := os.Getgid()
	fmt.Println(gid)
}

func Test_Getpid(t *testing.T) {
	pid := os.Getpid()
	fmt.Println(pid)
}

func TestExit(t *testing.T) {
	pid := os.Getpid()
	fmt.Println(pid)
	os.Exit(pid)
}
