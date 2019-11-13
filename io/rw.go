package io

import (
	"io"
)

type R struct {
	bys []byte
}

//func (r *R) Read(p []byte) (n int, err error) {
//	if r.bys == nil {
//		return 0, io.EOF
//	}
//	copy(p, r.bys)
//	r.bys = nil
//	return len(p), nil
//}

//mutil 使用
func (r *R) Read(p []byte) (n int, err error) {
	if r.bys == nil {
		return 0, io.EOF
	}

	index := 0
	for i, v := range p {
		if v == 0 {
			index = i
			break
		}
	}

	copy(p[index:], r.bys)
	r.bys = nil
	return 0, io.EOF
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
	if w.bys == nil {
		w.bys = make([]byte, 0)
	}
	w.bys = append(w.bys, p...)
	return len(p), nil
}

func (w *W) ReadFrom(r io.Reader) (n int64, err error) {
	rn, err := r.Read(w.bys)
	return int64(rn), err
}
