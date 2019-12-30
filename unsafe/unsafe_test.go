package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func Test_d1(t *testing.T) {
	var p float64 = 99
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p)))
}
