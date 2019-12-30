package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

//三种高性能的反射(针对不同数据类型)
func Test_offset(t *testing.T) {
	type TestObj struct {
		field1 string
	}
	struct_ := &TestObj{}
	field, _ := reflect.TypeOf(struct_).Elem().FieldByName("field1")
	field1Ptr := uintptr(unsafe.Pointer(struct_)) + field.Offset
	*((*string)(unsafe.Pointer(field1Ptr))) = "hello"
	fmt.Println(struct_)
}

func Test_interface(t *testing.T) {
	type TestObj struct {
		field1 string
	}
	struct_ := &TestObj{}
	structInter := (interface{})(struct_)
	type emptyInterface struct {
		typ  *struct{}
		word unsafe.Pointer
	}
	structPtr := (*emptyInterface)(unsafe.Pointer(&structInter)).word
	field, _ := reflect.TypeOf(structInter).Elem().FieldByName("field1")
	field1Ptr := uintptr(structPtr) + field.Offset
	*((*string)(unsafe.Pointer(field1Ptr))) = "hello"
	fmt.Println(struct_)
}

func Test_slice(t *testing.T) {
	slice := []string{"hello", "world"}
	type sliceHeader struct {
		Data unsafe.Pointer
		Len  int
		Cap  int
	}
	header := (*sliceHeader)(unsafe.Pointer(&slice))
	fmt.Println(header.Len)
	elementType := reflect.TypeOf(slice).Elem()
	secondElementPtr := uintptr(header.Data) + elementType.Size()
	*((*string)(unsafe.Pointer(secondElementPtr))) = "!!!"
	fmt.Println(slice)
}
