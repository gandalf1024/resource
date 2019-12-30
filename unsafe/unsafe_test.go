package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func Test_Sizeof(t *testing.T) {
	var p float64 = 99
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p)))
}

func Test_Offsetof(t *testing.T) {
	type student struct {
		name string
		age  int
		addr string
	}

	s1 := student{name: "zj", age: 11, addr: "123"}
	offname := unsafe.Offsetof(s1.name)
	offage := unsafe.Offsetof(s1.age)
	offaddr := unsafe.Offsetof(s1.addr)
	fmt.Println(offname) // 0
	fmt.Println(offage)  // 16
	fmt.Println(offaddr) // 24
}

func Test_Alignof(t *testing.T) {
	type student struct {
		name string
		age  int
		addr string
	}
	s1 := student{name: "zj", age: 11, addr: "123"}

	aliname := unsafe.Alignof(s1.name)
	aliage := unsafe.Alignof(s1.age)
	aliaddr := unsafe.Alignof(s1.addr)
	fmt.Println(aliname) // 8
	fmt.Println(aliage)  // 8
	fmt.Println(aliaddr) // 8
}
