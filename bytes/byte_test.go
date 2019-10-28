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
	by1 = append(by1, 49)
	by1 = append(by1, 50)
	by1 = append(by1, 51)

	flag1 := bytes.ContainsAny(by1, "2") //遍历i字符串每一项是 (字符编码)对应UTF8
	fmt.Println(flag1)
	flag2 := bytes.ContainsAny(by1, "1")
	fmt.Println(flag2)
	flag3 := bytes.ContainsAny(by1, "88")
	fmt.Println(flag3)
}

func Test_ContainsRune(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 49)
	by1 = append(by1, 50)
	by1 = append(by1, 51)
	a := rune('1')
	flag := bytes.ContainsRune(by1, a)
	fmt.Println(flag)
}

func Test_IndexByte(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	b := byte(11)
	index := bytes.IndexByte(by1, b) //调用汇编
	fmt.Println(index)
}

func Test_LastIndex(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 56)
	by1 = append(by1, 16)
	by1 = append(by1, 10) //--最后一个
	by1 = append(by1, 11)
	by1 = append(by1, 88)
	by1 = append(by1, 22)

	by2 := make([]byte, 0) //字串
	by2 = append(by2, 10)
	by2 = append(by2, 11)

	index := bytes.LastIndex(by1, by2) //调用Rabin-Karp算法
	fmt.Println(index)

}

func Test_LastIndex_Other(t *testing.T) {
	i := 1
	fmt.Println(i & 1)
	i = 2
	fmt.Println(i & 1)
	i = 3
	fmt.Println(i & 1)
	i = 4
	fmt.Println(i & 1)
}

func Test_LastIndexByte(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by1 = append(by1, 13)
	by1 = append(by1, 10)
	by1 = append(by1, 14)
	by1 = append(by1, 15)
	b := byte(10)
	index := bytes.LastIndexByte(by1, b)
	fmt.Println(index)

}

func Test_IndexRune(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	r := rune(11)
	index := bytes.IndexRune(by1, r)
	fmt.Println(index)

}

func Test_IndexAny(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 49)
	by1 = append(by1, 50)
	by1 = append(by1, 51)
	by1 = append(by1, 52)
	index := bytes.IndexAny(by1, "2")
	fmt.Println(index)
}

func Test_LastIndexAny(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 49)
	by1 = append(by1, 50)
	by1 = append(by1, 51)
	by1 = append(by1, 49)
	by1 = append(by1, 52)
	index := bytes.LastIndexAny(by1, "1") // 字符串 1 对应 utf8字符码 49
	fmt.Println(index)
}

func Test_SplitN(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10) // ---
	by1 = append(by1, 11) // ---
	by1 = append(by1, 12)
	by1 = append(by1, 13)
	by1 = append(by1, 10) // ---
	by1 = append(by1, 11) // ---
	by1 = append(by1, 14)
	by1 = append(by1, 15)
	by1 = append(by1, 10) // ---
	by1 = append(by1, 11) // ---
	by1 = append(by1, 16)

	by2 := make([]byte, 0)
	by2 = append(by2, 10)
	by2 = append(by2, 11)

	by3 := bytes.SplitN(by1, by2, 3) // n 为截出来几个数组
	for _, v := range by3 {
		fmt.Println(v)
	}
}

func Test_SplitAfterN(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by1 = append(by1, 13)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 14)
	by1 = append(by1, 15)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 16)
	by1 = append(by1, 17)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 18)
	by1 = append(by1, 19)

	by2 := make([]byte, 0)
	by2 = append(by2, 10)
	by2 = append(by2, 11)

	bys := bytes.SplitAfterN(by1, by2, 3) //after意思是  保留截取的数组by2
	for _, v := range bys {
		fmt.Println(v)
	}

}

func Test_Split(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10) // --
	by1 = append(by1, 11) // --
	by1 = append(by1, 12)
	by1 = append(by1, 13)
	by1 = append(by1, 10) // --
	by1 = append(by1, 11) // --
	by1 = append(by1, 14)
	by1 = append(by1, 15)
	by1 = append(by1, 10) // --
	by1 = append(by1, 11) // --
	by1 = append(by1, 16)

	by2 := make([]byte, 0)
	by2 = append(by2, 10)
	by2 = append(by2, 11)

	bys := bytes.Split(by1, by2)
	for _, v := range bys {
		fmt.Println(v)
	}

}

func Test_SplitAfter(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by1 = append(by1, 13)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 14)
	by1 = append(by1, 15)
	by1 = append(by1, 16)

	by2 := make([]byte, 0)
	by2 = append(by2, 10)
	by2 = append(by2, 11)

	bys := bytes.SplitAfter(by1, by2)
	for _, v := range bys {
		fmt.Println(v)
	}

}

func Test_Fields(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 0)
	by1 = append(by1, 1)
	by1 = append(by1, 2)
	by1 = append(by1, 3)
	by1 = append(by1, 4)
	by1 = append(by1, 5)
	by1 = append(by1, 6)
	by1 = append(by1, 7)
	by1 = append(by1, 8)
	by1 = append(by1, 9)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by1 = append(by1, 13)
	by1 = append(by1, 200)

	bys := bytes.Fields(by1)
	for _, v := range bys {
		fmt.Println(v)
	}

}

func Test_Fields_Other(t *testing.T) {
	var a = uint8(11)
	var b = uint8(22)
	a |= b //或等运算
	println(a)
	println(b)
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
