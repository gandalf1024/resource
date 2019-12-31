package slice

import (
	"fmt"
	"testing"
)

func Test_slice(t *testing.T) {
	var array [10]int
	var slice = array[5:6]
	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
}

func AddElement(slice []int, e int) []int {
	return append(slice, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e, e)
}

func Test_slice2(t *testing.T) {
	var slice []int
	slice = append(slice, 1, 2, 3)
	newSlice := AddElement(slice, 4)

	fmt.Println(slice, "----1")
	fmt.Println(newSlice, "----2")
	fmt.Println(&slice[0] == &newSlice[0]) // 当前添加的元素没有导致底层数组扩容时使用同一数组所以相等,但是添加过多元素导致底层数组扩展就不相等
}

func Test_slice3(t *testing.T) {
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(pollorder) = ", len(pollorder))
	fmt.Println("cap(pollorder) = ", cap(pollorder))
	fmt.Println("len(lockorder) = ", len(lockorder))
	fmt.Println("cap(lockorder) = ", cap(lockorder))
}
