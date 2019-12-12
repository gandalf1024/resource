package path

import (
	"fmt"
	"path"
	"testing"
)

func Test_Clean(t *testing.T) {
	p := path.Clean("//a//b//c//")
	fmt.Println(p)

	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
	}
	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

}

func Test_Join(t *testing.T) {
	p := path.Join("a", "b", "c")
	fmt.Println(p) // a/b/c
}

func Test_Split(t *testing.T) {
	dir, file := path.Split("a/s/c/v/aa.txt")
	fmt.Println("dir:=", dir)
	fmt.Println("file:=", file)
}

func Test_Ext(t *testing.T) {
	str := path.Ext("a/b/c/cc.txt") // 取格式
	fmt.Println(str)
}

func Test_Base(t *testing.T) {
	b := path.Base("/a/b/c/v/cc.txt")
	fmt.Println(b)
}

func Test_IsAbs(t *testing.T) {
	is := path.IsAbs("./data")
	fmt.Println(is)
	is = path.IsAbs("D:\\code\\golang\\src\\resource\\path\\data\\data.txt")
	fmt.Println(is)
}

func Test_Dir(t *testing.T) {
	fmt.Println(path.Dir("/a/b/c"))
}
