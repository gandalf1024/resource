package bytes

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func Test_NewBuffer(t *testing.T) {

	file, err := os.Open("./data/aa.txt")
	if err != nil {
		panic(err)
	}
	bys, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(bys)

	//fmt.Println(string(buffer.Bytes()))
	//fmt.Println(buffer.String())
	//fmt.Println(buffer.Len())
	//fmt.Println(buffer.Cap())
	buffer.Truncate(1)
	buffer.Write([]byte("---------------------11111111---------\n"))
	//fmt.Println(buffer.String())
	buffer.Write([]byte("457634765536873567346573547654367654\n"))
	//fmt.Println(buffer.String())
	buffer.WriteString("3546346575467456775467\n")
	fmt.Println(buffer.String())
	fmt.Println(buffer.Cap())
	buffer.Grow(5000)
	fmt.Println(buffer.Cap())
	buffer.WriteRune('或')

	n, _ := buffer.Read([]byte("3252ewfedgfsadg"))
	fmt.Println(n)
	buffer.ReadRune()
	err = buffer.UnreadRune() //回退到ReadRune之前
	if err != nil {
		panic(err)
	}

	line, err := buffer.ReadBytes('-')
	if err != nil {
		panic(err)
	}
	fmt.Println(string(line))

}

func Test_NewBuffer_Demo(t *testing.T) {
	type readOp int8

	const (
		opRead      readOp = -1
		opInvalid   readOp = 0
		opReadRune1 readOp = 1
		opReadRune2 readOp = 2
		opReadRune3 readOp = 3
		opReadRune4 readOp = 4
	)

	fmt.Println(readOp(2))

	indx := bytes.IndexByte([]byte("1212"), 'j')
	fmt.Println(indx)

	fmt.Println(string(45))

	var line []byte
	line = append(line, 11)
	fmt.Println(line)

}

type R struct {
	bys []byte
}

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
	return len(r.bys), io.EOF
}

func Test_ReadFrom(t *testing.T) {
	r := &R{}
	r.bys = []byte("99999")

	bf := bytes.NewBuffer([]byte("111111111"))
	n, err := bf.ReadFrom(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func Test_ReadFrom_file(t *testing.T) {
	f, err := os.Open("./data/aa.txt")
	if err != nil {
		panic(err)
	}
	bf := bytes.NewBuffer([]byte("123456789"))
	n, err := bf.ReadFrom(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
