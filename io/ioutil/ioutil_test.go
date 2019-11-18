package ioutil

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_Reader(t *testing.T) {
	r := &R{}
	r.bys = []byte("13452364")

	bys, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bys))
}

func Test_Sli(t *testing.T) {
	bys := []byte("123456789")
	Change(bys[:2])
	fmt.Println(string(bys))
	Change2(bys)
	fmt.Println(string(bys))

}

func Change(bb []byte) {
	copy(bb[2:], "9687461561")
}

func Change2(bb []byte) {
	copy(bb[2:], "9687461561")
}
