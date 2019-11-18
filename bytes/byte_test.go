package bytes

import (
	"bytes"
	"fmt"
	"testing"
)

var compareTests = []struct {
	a, b []byte
	i    int
}{
	{[]byte(""), []byte(""), 0},
	{[]byte("a"), []byte(""), 1},
	{[]byte(""), []byte("a"), -1},
	{[]byte("abc"), []byte("abc"), 0},
	{[]byte("abd"), []byte("abc"), 1},
	{[]byte("abc"), []byte("abd"), -1},
	{[]byte("ab"), []byte("abc"), -1},
	{[]byte("abc"), []byte("ab"), 1},
	{[]byte("x"), []byte("ab"), 1},
	{[]byte("ab"), []byte("x"), -1},
	{[]byte("x"), []byte("a"), 1},
	{[]byte("b"), []byte("x"), -1},
	// test runtime·memeq's chunked implementation
	{[]byte("abcdefgh"), []byte("abcdefgh"), 0},
	{[]byte("abcdefghi"), []byte("abcdefghi"), 0},
	{[]byte("abcdefghi"), []byte("abcdefghj"), -1},
	{[]byte("abcdefghj"), []byte("abcdefghi"), 1},
	// nil tests
	{nil, nil, 0},
	{[]byte(""), nil, 0},
	{nil, []byte(""), 0},
	{[]byte("a"), nil, 1},
	{nil, []byte("a"), -1},
}

func Test_Equal(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by2 := make([]byte, 0)
	by2 = append(by2, 10)

	flag := bytes.Equal(by1, by2) //先转换为string然后判断字符串是否相等     （原理string的底层是byte数组）
	fmt.Println(flag)
}

func Test_Equal_Off(t *testing.T) {

	allocs := testing.AllocsPerRun(10, func() { // 调用10次
		for _, tt := range compareTests {
			eql := bytes.Equal(tt.a, tt.b)
			if eql != (tt.i == 0) {
				t.Errorf(`Equal(%q, %q) = %v`, tt.a, tt.b, eql)
			}
		}
	})

	if allocs > 0 {
		t.Errorf("Equal allocated %v times", allocs)
	}

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

// 9 10 11 12 13 32 133 160  这些数字会切数组
// 原因是isSpace 里面的判断
func Test_Fields(t *testing.T) {
	by1 := make([]byte, 0)

	by1 = append(by1, 0)
	by1 = append(by1, 1)
	by1 = append(by1, 2)
	by1 = append(by1, 3)
	by1 = append(by1, 4)
	by1 = append(by1, 5)
	by1 = append(by1, 6)
	//by1 = append(by1, 7)
	//by1 = append(by1, 8)
	//by1 = append(by1, 9)
	//by1 = append(by1, 10)
	//by1 = append(by1, 11)
	//by1 = append(by1, ' ')
	//by1 = append(by1, 12)
	//by1 = append(by1, 13)
	//by1 = append(by1, 200)

	bys := bytes.Fields(by1)
	for _, v := range bys {
		fmt.Println(v)
	}

	//bys = bytes.FieldsFunc(by1, unicode.IsSpace)
	//for _, v := range bys {
	//	fmt.Println(v)
	//}

}

func Test_Fields_Other(t *testing.T) {
	var a = uint8(11)
	var b = uint8(22)
	a |= b //或等运算
	println(a)
	println(b)

	var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

	fmt.Println(asciiSpace[0])

	n := 0
	wasSpace := 1
	isSpace := 0
	n += wasSpace & ^isSpace
	fmt.Println(n)
	fmt.Println(^isSpace)
	fmt.Println(wasSpace & wasSpace)
	fmt.Println(wasSpace & ^isSpace)
	fmt.Println(0x80)
	fmt.Println(0x85, 0xA0)

	r := 9
	if uint32(r) <= '\u00FF' {
		switch r {
		case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
			fmt.Println('\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0)
		}
		fmt.Println(r)
	}

}

func Test_FieldsFunc(t *testing.T) {
	by1 := make([]byte, 0)

	by1 = append(by1, 0)
	by1 = append(by1, 1)
	by1 = append(by1, 2)
	by1 = append(by1, 3)
	by1 = append(by1, 4)
	by1 = append(by1, 5)
	by1 = append(by1, 6)

	bys := bytes.FieldsFunc(by1, func(r rune) bool {
		if r == 1 || r == 3 {
			return true
		}
		return false
	})

	for _, v := range bys {
		fmt.Println(v)
	}

}

func Test_Join(t *testing.T) {
	by1 := make([][]byte, 0)
	by11 := make([]byte, 0)
	by11 = append(by11, 0)
	by11 = append(by11, 1)
	by11 = append(by11, 2)
	by11 = append(by11, 3)
	by11 = append(by11, 4)
	by11 = append(by11, 5)
	by11 = append(by11, 6)
	by11 = append(by11, 116)

	by12 := make([]byte, 0)
	by12 = append(by12, 0)
	by12 = append(by12, 1)
	by12 = append(by12, 2)
	by12 = append(by12, 3)
	by12 = append(by12, 4)
	by12 = append(by12, 5)
	by12 = append(by12, 6)
	by12 = append(by12, 116)

	by1 = append(by1, by11)
	by1 = append(by1, by12)

	by2 := make([]byte, 0) //中间分割的切片
	by2 = append(by2, 25)
	by2 = append(by2, 28)
	by2 = append(by2, 49)

	// 连接二维切片中的所有数据为一个切片，使用by2分割切片间的数据
	bys := bytes.Join(by1, by2)
	fmt.Println(bys)

}

func Test_HasPrefix(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 34)
	by1 = append(by1, 25)
	by1 = append(by1, 56)

	by2 := make([]byte, 0)
	by2 = append(by2, 34)

	res := bytes.HasPrefix(by1, by2)
	fmt.Println(res)
}

func Test_HasPrefix_Demo(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 34)
	by1 = append(by1, 25)
	by1 = append(by1, 56)

	by2 := make([]byte, 0)
	by2 = append(by2, 34)

	res := bytes.Equal(by1[:1], by2)
	fmt.Println(res)
}

func Test_Join_Demo(t *testing.T) {
	by2 := make([]byte, 0)
	by2 = append(by2, 25)
	by2 = append(by2, 28)
	by2 = append(by2, 49)

	by3 := make([]byte, 0)
	by3 = append(by3, 99)

	bp := copy(by2, by3)
	fmt.Println(bp)
	fmt.Println(by2)
	bp += copy(by2[bp:], by3)
	fmt.Println(bp)
	fmt.Println(by2)

}

func Test_HasSuffix(t *testing.T) {
	by2 := make([]byte, 0)
	by2 = append(by2, 25)
	by2 = append(by2, 28)
	by2 = append(by2, 49)

	by3 := make([]byte, 0)
	by3 = append(by3, 49)

	res := bytes.HasSuffix(by2, by3)
	fmt.Println(res)
}

func Test_Map(t *testing.T) {
	by2 := make([]byte, 0)
	by2 = append(by2, 1)
	by2 = append(by2, 2)
	by2 = append(by2, 3)
	by2 = append(by2, 4)
	by2 = append(by2, 5)
	by2 = append(by2, 6)

	f := func(r rune) rune {
		if r == 3 {
			return 99
		}
		return r
	}

	bys := bytes.Map(f, by2) //替换byte数组中的指定值

	fmt.Println(bys)
}

func Test_Repeat(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 1)
	by1 = append(by1, 2)
	by1 = append(by1, 3)
	by1 = append(by1, 4)
	by1 = append(by1, 5)

	bys := bytes.Repeat(by1, 3) //复制切片  count次数
	fmt.Println(bys)
}

func Test_ToUpper(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 'a')
	by1 = append(by1, 'b')
	by1 = append(by1, 'c')
	by1 = append(by1, 'd')
	by1 = append(by1, 'e')
	by1 = append(by1, 100)

	bys := bytes.ToUpper(by1)
	for _, v := range bys {
		fmt.Println(string(v))
	}

}

func Test_ToLower(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 'A')
	by1 = append(by1, 'B')
	by1 = append(by1, 'C')
	by1 = append(by1, 'D')
	by1 = append(by1, 'E')
	by1 = append(by1, 100)

	bys := bytes.ToLower(by1)
	for _, v := range bys {
		fmt.Println(string(v))
	}

}

func Test_ToTitle(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 'a')
	by1 = append(by1, 'b')
	by1 = append(by1, 'c')
	by1 = append(by1, 'd')
	by1 = append(by1, 'e')
	by1 = append(by1, ' ')
	by1 = append(by1, 'q')
	by1 = append(by1, 'w')
	by1 = append(by1, 't')

	bys := bytes.Title(by1) //首字符大写
	for _, v := range bys {
		fmt.Println(string(v))
	}
}

func Test_TrimLeftFunc(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, '1')
	by1 = append(by1, '1')
	by1 = append(by1, 'a')
	by1 = append(by1, 'b')
	by1 = append(by1, 'c')
	by1 = append(by1, 'd')

	bys := bytes.TrimLeftFunc(by1, func(r rune) bool {
		if r == '1' {
			return true
		}
		return false
	})

	for _, v := range bys {
		fmt.Println(string(v))
	}
}

func Test_TrimRightFunc(t *testing.T) {
	by1 := make([]byte, 0)

	by1 = append(by1, 'a')
	by1 = append(by1, 'a')
	by1 = append(by1, 'b')
	by1 = append(by1, 'c')
	by1 = append(by1, 'd')
	by1 = append(by1, '1')
	by1 = append(by1, '1')

	bys := bytes.TrimRightFunc(by1, func(r rune) bool {
		if r == '1' {
			return true
		}
		return false
	})

	for _, v := range bys {
		fmt.Println(string(v))
	}

}

func Test_TrimFunc(t *testing.T) {
	by1 := make([]byte, 0)

	by1 = append(by1, '1')
	by1 = append(by1, '1')
	by1 = append(by1, 'a')
	by1 = append(by1, 'a')
	by1 = append(by1, 'b')
	by1 = append(by1, 'c')
	by1 = append(by1, 'd')
	by1 = append(by1, '1')
	by1 = append(by1, '1')

	bys := bytes.TrimFunc(by1, func(r rune) bool {
		if r == '1' {
			return true
		}
		return false
	})

	for _, v := range bys {
		fmt.Println(string(v))
	}

}

func Test_TrimPrefix(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 'a')
	by1 = append(by1, 'a')
	by1 = append(by1, 'b')
	by1 = append(by1, 'c')
	by1 = append(by1, 'd')
	by1 = append(by1, '1')
	by1 = append(by1, '1')

	by2 := make([]byte, 0)
	by2 = append(by2, 'a')
	by2 = append(by2, 'a')
	by2 = append(by2, 'b')

	bys := bytes.TrimPrefix(by1, by2)
	for _, v := range bys {
		fmt.Println(string(v))
	}

}

func Test_TrimSuffix(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 'a')
	by1 = append(by1, 'a')
	by1 = append(by1, 'b')
	by1 = append(by1, 'c')
	by1 = append(by1, 'd')
	by1 = append(by1, '1')
	by1 = append(by1, '1')

	by2 := make([]byte, 0)
	by2 = append(by2, 'd')
	by2 = append(by2, '1')
	by2 = append(by2, '1')

	bys := bytes.TrimSuffix(by1, by2)

	for _, v := range bys {
		fmt.Println(string(v))
	}

}

func Test_IndexFunc(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 1)
	by1 = append(by1, 2)
	by1 = append(by1, 3)
	by1 = append(by1, 4)

	res := bytes.IndexFunc(by1, func(r rune) bool {
		if r == 3 {
			return true
		}
		return false
	})

	fmt.Println(res)
}

func Test_LastIndexFunc(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 1)
	by1 = append(by1, 2)
	by1 = append(by1, 3)
	by1 = append(by1, 4)

	index := bytes.LastIndexFunc(by1, func(r rune) bool {
		if r == 4 {
			return true
		}
		return false
	})
	fmt.Println(index)
}

func Test_Trim(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, '1')
	by1 = append(by1, '2')
	by1 = append(by1, 3)
	by1 = append(by1, 4)

	bys := bytes.Trim(by1, "123")
	fmt.Println(bys)
}

func Test_TrimLeft(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, '1')
	by1 = append(by1, '2')
	by1 = append(by1, 3)
	by1 = append(by1, 4)

	bys := bytes.TrimLeft(by1, "123")

	for _, v := range bys {
		fmt.Println(v)
	}
}

func Test_TrimRight(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 4)
	by1 = append(by1, 3)
	by1 = append(by1, '2')
	by1 = append(by1, '1')

	bys := bytes.TrimRight(by1, "123")

	for _, v := range bys {
		fmt.Println(v)
	}
}

func Test_TrimSpace(t *testing.T) {
	fmt.Println("    asdfghjkl    ")
	bys := bytes.TrimSpace([]byte("    asdfghjkl    "))
	fmt.Println(string(bys))

	by1 := make([]byte, 0)
	by1 = append(by1, ' ')
	by1 = append(by1, ' ')
	by1 = append(by1, '2')
	by1 = append(by1, '3')
	by1 = append(by1, '4')
	by1 = append(by1, '5')
	by1 = append(by1, ' ')
	by1 = append(by1, ' ')

	bys2 := bytes.TrimSpace(by1)
	fmt.Println(bys2)

}

func Test_Runes(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 1)
	by1 = append(by1, 2)
	by1 = append(by1, 3)
	by1 = append(by1, 4)

	res := bytes.Runes(by1)

	for _, v := range res {
		fmt.Println(v)
	}

}

func Test_Replace(t *testing.T) {
	s := make([]byte, 0) //源
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 5)
	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 8)
	s = append(s, 3)
	s = append(s, 4)
	s = append(s, 5)

	old := make([]byte, 0) //源切片中的 34
	old = append(old, 3)
	old = append(old, 4)

	news := make([]byte, 0) // 替换为 7890
	news = append(news, 7)
	news = append(news, 8)
	news = append(news, 9)
	news = append(news, 0)

	bys := bytes.Replace(s, old, news, 2) //n 替换几个   -1 全替换
	for _, v := range bys {
		fmt.Println(v)
	}

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

func Test_EqualFold(t *testing.T) {
	by1 := make([]byte, 0)
	by1 = append(by1, 10)
	by1 = append(by1, 11)
	by1 = append(by1, 12)
	by1 = append(by1, 14)

	by2 := make([]byte, 0)
	by2 = append(by2, 10)
	by2 = append(by2, 11)
	by2 = append(by2, 12)
	by2 = append(by2, 13)

	res := bytes.EqualFold(by1, by2)
	fmt.Println(res)

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

	indx := bytes.Index([]byte("asdfghjkl;"), []byte("g"))
	fmt.Println(indx)

}

func Test_Index_Demo(t *testing.T) {
	fmt.Println(indexRabinKarp([]byte("asdfghjkl"), []byte("gh")))
}

func indexRabinKarp(s, sep []byte) int {
	hashsep, pow := hashStr(sep)
	n := len(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashsep && bytes.Equal(s[:n], sep) {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashsep && bytes.Equal(s[i-n:i], sep) {
			return i - n
		}
	}
	return -1
}

const primeRK = 16777619

func hashStr(sep []byte) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
