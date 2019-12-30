package builtin

import (
	"fmt"
	"log"
	"math/big"
	"testing"
	"unsafe"
)

//默认值
func Test_f1(t *testing.T) {
	var a bool

	var b int8
	var c int16
	var d int32
	var e int64
	var f uint8
	var g uint16
	var h uint32
	var i uint64
	var j uintptr

	var k complex64
	var l complex128

	var m byte
	var n rune

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)

}

//赋值范围
func Test_f2(t *testing.T) {
	var a bool = true
	var aa bool = false

	var b int8 = 127
	var bb int8 = -128
	var c int16 = -32768
	var cc int16 = 32767
	var d int32 = -2147483648
	var dd int32 = 2147483647
	var e int64 = -9223372036854775808
	var ee int64 = 9223372036854775807
	var f uint8 = 255
	var g uint16 = 65535
	var h uint32 = 4294967295
	var i uint64 = 18446744073709551615
	var j uintptr //足够存储一个指针地址

	var k complex64  //8字节
	var l complex128 //16字节

	var m byte = 255         //byte = uint8
	var n rune = -2147483648 //rune = int32
	var nn rune = 2147483647 //rune = int32

	fmt.Println("1字节等于8位(0000 0000)")
	fmt.Println(a, "==bool==", unsafe.Sizeof(a))
	fmt.Println(a, "==bool==", unsafe.Sizeof(aa))
	fmt.Println(b, "==int8==", unsafe.Sizeof(b))
	fmt.Println(bb, "==int8==", unsafe.Sizeof(bb))
	fmt.Println(c, "==int16==", unsafe.Sizeof(c))
	fmt.Println(cc, "==int16==", unsafe.Sizeof(cc))
	fmt.Println(d, "==int32==", unsafe.Sizeof(d))
	fmt.Println(dd, "==int32==", unsafe.Sizeof(dd))
	fmt.Println(e, "==int64==", unsafe.Sizeof(e))
	fmt.Println(ee, "==int64==", unsafe.Sizeof(ee))
	fmt.Println(f, "==uint8==", unsafe.Sizeof(f))
	fmt.Println(g, "==uint16==", unsafe.Sizeof(g))
	fmt.Println(h, "==uint32==", unsafe.Sizeof(h))
	fmt.Println(i, "==uint64==", unsafe.Sizeof(i))
	fmt.Println(j, "==uintptr==", unsafe.Sizeof(j))
	fmt.Println(k, "==complex64==", unsafe.Sizeof(k))
	fmt.Println(l, "==complex128==", unsafe.Sizeof(l))
	fmt.Println(m, "==byte==", unsafe.Sizeof(m))
	fmt.Println(n, "==rune==", unsafe.Sizeof(n))
	fmt.Println(nn, "==rune==", unsafe.Sizeof(nn))
}

//计算范围
func Test_f3(t *testing.T) {
	var i uint64 = 1
	fmt.Println("int:")
	fmt.Println(i << 8 / 2)
	fmt.Println(i << 16 / 2)
	fmt.Println(i << 32 / 2)
	//左移少一位相当于除以2
	fmt.Println(i << 63)

	fmt.Println("uint:")
	fmt.Println(i << 8)
	fmt.Println(i << 16)
	fmt.Println(i << 32)
	//左移少一位相当于除以2
	fmt.Println(i << 63)

	//大数计算
	ivv := new(big.Int)
	_, err := fmt.Sscan("9223372036854775808", ivv)

	ivv2 := new(big.Int)
	_, err = fmt.Sscan("2", ivv2)

	result := new(big.Int)

	mul := result.Mul(ivv, ivv2)
	fmt.Println(mul)

	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(i)
	}
}
