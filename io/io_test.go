package io

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type S struct {
	data string
}

func (str *S) WriteString(s string) (n int, err error) {
	str.data += s
	return len(s) * 10, nil
}

func (str *S) Write(p []byte) (n int, err error) {
	str.data += string(p)
	return len(p) * 5, nil
}

func Test_StringWriter(t *testing.T) {
	str := &S{}
	n, err := io.WriteString(str, "你hi撒旦广泛的事故发生的")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

type R struct {
	bys []byte
}

func (r *R) Read(p []byte) (n int, err error) {
	if r.bys == nil {
		return 0, io.EOF
	}
	copy(p, r.bys)
	r.bys = nil
	return len(p), nil
}

func (r *R) WriteTo(w io.Writer) (n int64, err error) {
	ncount, err := w.Write(r.bys)
	return int64(ncount), err
}

func (r *R) ReadAt(p []byte, off int64) (n int, err error) {
	if r.bys == nil {
		return 0, io.EOF
	}
	copy(p, r.bys[off:])
	r.bys = nil
	return len(p), nil
}

type W struct {
	bys []byte
}

func (w *W) Write(p []byte) (n int, err error) {
	w.bys = append(w.bys, p...)
	return len(p), nil
}

func (w *W) ReadFrom(r io.Reader) (n int64, err error) {
	rn, err := r.Read(w.bys)
	return int64(rn), err
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

func Test_ReaderAt(t *testing.T) {
	r := &R{}
	var ra io.ReaderAt = r
	sectionReader := io.NewSectionReader(ra, 1, 1)

	bys := make([]byte, 100)
	n, err := sectionReader.Read(bys)
	if err != nil {
		panic(err)
	}

	fmt.Println(n)
	fmt.Println(string(bys))
}
