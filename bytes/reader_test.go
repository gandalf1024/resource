package bytes

import (
	"bytes"
	"fmt"
	"testing"
)

var r = bytes.NewReader([]byte("464546545654546"))

func Test_NewReader(t *testing.T) {
	reder := bytes.NewReader([]byte("464546545654546"))
	fmt.Println(reder)
}

func Test_Len(t *testing.T) {
	fmt.Println(r.Len())
}

func Test_Size(t *testing.T) {
	fmt.Println(r.Size())
}

func Test_Read(t *testing.T) {
	var b = make([]byte, 100)
	i, err := r.Read(b) //把数据读取到b中
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

}

func Test_ReadRune(t *testing.T) {
	ch, size, err := r.ReadRune()
	if err != nil {
		panic(err)
	}

	fmt.Println(ch, size)

}

func Test_Seek(t *testing.T) {
	idx, err := r.Seek(8000, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(idx)
}

func Test_WriteTo(t *testing.T) {
	v := r
	var b bytes.Buffer
	n, _ := v.WriteTo(&b)
	fmt.Println(n)

}

func Test_CopyN(t *testing.T) {

}
