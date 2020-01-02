package slice

import (
	"fmt"
	"testing"
	"time"
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

	// Tips:
	// 使用append()向切片追加元素时有可能触发扩容，扩容后将会生成新的切片
}

func Test_ArrayToSlice(t *testing.T) {
	arr := [3]int64{11, 22, 33}
	sli := arr[:1]

	fmt.Println(sli)
	fmt.Println("len:", len(sli))
	fmt.Println("cap:", cap(sli))
	// Tips:
	// 切片共用数组空间
}

//测试切片扩容规则 (每个版本不一样)
func Test_Expansion(t *testing.T) {
	sli := make([]int, 0, 0)

	capNum := 0

	index := 1
	for index < 100 {
		sli = append(sli, 50)
		if cap(sli) > capNum {
			fmt.Println("index:", index)
			fmt.Println("len:", len(sli))
			fmt.Println("cap:", cap(sli))
			index++
			time.Sleep(time.Millisecond * 50)
			capNum = cap(sli)
			fmt.Println("---------------------------------------------->>>>>>>>>")
		}
	}

	// Tips:
	//如果原Slice容量小于1024，则新Slice容量将扩大为原来的2倍；
	//如果原Slice容量大于等于1024，则新Slice容量将扩大为原来的1.25倍；
}

//测试同时添加和copy
func Test_copy(t *testing.T) {
	sli1 := make([]int, 0, 0)
	sli2 := make([]int, 0, 0)

	go func() {
		time.Sleep(time.Millisecond)
		index := 1
		for index < 1000 {
			sli1 = append(sli1, 50)
			if index%2 == 0 {
				sli2 = append(sli2, 60)
			}
			index++
		}
	}()

	time.Sleep(time.Millisecond)
	copy(sli1, sli2)

	fmt.Println(sli1)
	fmt.Println(len(sli1))

	// Tips:
	// copy不是并发安全的
}
