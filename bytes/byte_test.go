package d_02

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_Equal(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by2 := make([]byte, 0)
	by2 = append(by2, 10)

	flag := bytes.Equal(by1, by2) //先转换为string然后判断字符串是否相等     （原理string的底层是byte数组）
	fmt.Println(flag)
}

func Test_Compare(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by1 = append(by1, 13)
	by2 := make([]byte, 0)
	by2 = append(by2, 11)

	index := bytes.Compare(by2, by1) //底层调用汇编
	fmt.Println(index)
}

func Test_Count(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by2 := make([]byte, 0)
	by2 = append(by2, 20)

	count := bytes.Count(by1, by2)
	fmt.Println(count)
}

func Test_Contains(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by2 := make([]byte, 0)
	by2 = append(by2, 10)
	by2 = append(by2, 11)

	flag := bytes.Contains(by1, by2)
	fmt.Println(flag)

}

func Test_ContainsAny(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)

	flag := bytes.ContainsAny(by1, "10")
	fmt.Println(flag)

}

func Test_Index(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 12)
	by1 = append(by1, 45)
	by2 := make([]byte, 0)
	by2 = append(by2, 32)
	by2 = append(by2, 67)

	in := bytes.Index(by1, by2)
	fmt.Println(in)

	fmt.Println("----------------")

	by3 := make([]byte, 0)
	by3 = append(by3, 10)
	by3 = append(by3, 20)
	by4 := make([]byte, 0)
	//by4 = append(by4, 10)
	by4 = append(by4, 20)

	in2 := bytes.Index(by3, by4)
	fmt.Println(in2)

}

func Test_ReplaceAll(t *testing.T) {
	by1 := make([]byte, 0) //替换的切片
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by2 := make([]byte, 0)
	by2 = append(by2, 20)
	//by2 = append(by2, 21)
	//by2 = append(by2, 22)
	by3 := make([]byte, 0)
	by3 = append(by3, 30)
	by3 = append(by3, 20) //相同点开始替换
	by3 = append(by3, 32)

	by4 := bytes.ReplaceAll(by3, by2, by1) //我有by3   包含by2为开始点替换为by1
	fmt.Println(by4)
	fmt.Println(by3)
	fmt.Println(by2)
	fmt.Println(by1)

}
