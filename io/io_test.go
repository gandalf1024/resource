package io

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_StringWriter(t *testing.T) {
	w := &W{}
	n, err := io.WriteString(w, "你hi撒旦广泛的事故发生的")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func Test_ReadAtLeast(t *testing.T) {
	r := &R{}
	r.bys = append(r.bys, 22)
	n, err := io.ReadAtLeast(r, []byte("11212312"), 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

	fmt.Println(len(r.bys))

}

func Test_ReadFull(t *testing.T) {
	r := &R{}
	n, err := io.ReadFull(r, []byte("123456"))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(len(r.bys))
}

func Test_CopyN(t *testing.T) {
	w := new(W)
	r := new(R)
	r.bys = []byte("456")
	w.bys = []byte("123")

	n, err := io.CopyN(w, r, 3) //n:可读取的字节数
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(w.bys))
	fmt.Println(string(r.bys))

	var s interface{} = r
	if wt, ok := s.(io.WriterTo); ok { // 结构体转接口  需要转interfaces
		fmt.Println(wt)
	}

	s = w
	if rt, ok := s.(io.ReaderFrom); ok {
		fmt.Println(rt)
	}

	_ = s
}

func Test_CopyN_Demo(t *testing.T) {
	w := new(W)
	r := new(R)
	wn, err := copyBuffer(w, r)
	if err != nil {
		panic(err)
	}
	fmt.Println(wn)
}

func copyBuffer(dst io.Writer, src io.Reader) (written int64, err error) {
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	return written, err
}

func Test_Copy(t *testing.T) {
	w := new(W)
	r := new(R)
	w.bys = []byte("123")
	r.bys = []byte("456")

	wn, err := io.Copy(w, r) // 把r copy到w
	if err != nil {
		panic(err)
	}
	fmt.Println(wn)
	fmt.Println(string(w.bys))
	fmt.Println(string(r.bys))
}

func Test_Copy_Off(t *testing.T) {
	type Buffer struct {
		bytes.Buffer
		io.ReaderFrom
		io.WriterTo
	}

	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		t.Errorf("Copy did not work properly")
	}

}

func Test_CopyBuffer(t *testing.T) {
	w := new(W)
	r := new(R)
	w.bys = []byte("123")
	r.bys = []byte("456")
	wn, err := io.CopyBuffer(w, r, []byte("qwertyui"))
	if err != nil {
		panic(err)
	}
	fmt.Println(wn)
	fmt.Println(string(w.bys))
	fmt.Println(string(r.bys))
}

func Test_CopyBuffer_Demo(t *testing.T) {
	f, err := os.Open("D:\\code\\golang\\src\\resource\\temp\\03\\main.go")
	if err != nil {
		panic(err)
	}
	bys := make([]byte, 10)

	n, err := f.Read(bys)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(bys))

}

func Test_LimitReader(t *testing.T) {
	r := &R{}
	r.bys = []byte("12345678910111213141516171819")
	reader := io.LimitReader(r, 10)

	bys := make([]byte, 100)

	n, err := reader.Read(bys)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(string(bys))

}

func Test_NewSectionReader_Read(t *testing.T) {
	r := &R{}
	r.bys = []byte("123456789101112")
	var ra io.ReaderAt = r
	sectionReader := io.NewSectionReader(ra, 2, 3) //从索引地址 2 开始  读取3个

	bys := make([]byte, 100)
	n, err := sectionReader.Read(bys)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(string(bys))
}

func Test_NewSectionReader_Seek(t *testing.T) {
	r := &R{}
	r.bys = []byte("123456789101112")
	var ra io.ReaderAt = r
	sectionReader := io.NewSectionReader(ra, 2, 3) //从索引地址 2 开始  读取3个

	n, err := sectionReader.Seek(2, 2) // 移动开始点
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
}

func Test_NewSectionReader_ReadAt(t *testing.T) {
	r := &R{}
	r.bys = []byte("123456789101112435678698657896789")
	var ra io.ReaderAt = r
	sectionReader := io.NewSectionReader(ra, 1, 5) //从索引地址 1 开始  读取2个

	bys := make([]byte, 100)
	n, err := sectionReader.ReadAt(bys, 3)
	if err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(n)
}

func Test_NewSectionReader_Size(t *testing.T) {
	r := &R{}
	r.bys = []byte("123456789101112435678698657896789")
	sectionReader := io.NewSectionReader(r, 3, 8) // size的值和n永远一致
	size := sectionReader.Size()
	fmt.Println(size)
}

func Test_TeeReader(t *testing.T) {
	r := &R{}
	w := &W{}

	r.bys = []byte("123456789")

	reader := io.TeeReader(r, w)
	bys := make([]byte, 100)
	n, err := reader.Read(bys)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(r.bys))
	fmt.Println(string(w.bys))
}
