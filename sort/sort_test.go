package sort

import (
	"fmt"
	"sort"
	"testing"
)

type num_sort struct {
	num []int
}

func (num_sort *num_sort) Len() int {
	return len(num_sort.num)
}

func (num_sort *num_sort) Less(i, j int) bool {
	return num_sort.num[i] < num_sort.num[j] // 决定顺序
}

func (num_sort *num_sort) Swap(i, j int) {
	num_sort.num[i], num_sort.num[j] = num_sort.num[j], num_sort.num[i]
}

func Test_Interface(t *testing.T) {
	num_sort := &num_sort{num: make([]int, 0)}
	num_sort.num = append(num_sort.num, 5)
	num_sort.num = append(num_sort.num, 3)
	num_sort.num = append(num_sort.num, 7)
	num_sort.num = append(num_sort.num, 2)
	num_sort.num = append(num_sort.num, 34)
	num_sort.num = append(num_sort.num, 12)
	num_sort.num = append(num_sort.num, 23)
	num_sort.num = append(num_sort.num, 54)
	num_sort.num = append(num_sort.num, 16)
	num_sort.num = append(num_sort.num, 28)
	num_sort.num = append(num_sort.num, 67)
	num_sort.num = append(num_sort.num, 23)
	num_sort.num = append(num_sort.num, 14)

	sort.Sort(num_sort) //
	//Sort(num_sort, len(num_sort.num))

	for _, v := range num_sort.num {
		fmt.Println(v)
	}
}

func Sort(data sort.Interface, b int) {
	for i := 1; i < b; i++ {
		for j := i; j > j-1 && data.Less(j, j-1); j-- { // 相当于冒泡排序
			data.Swap(j, j-1)
		}
	}
}

func Test_Interface_demo(t *testing.T) {
	i := 40
	i >>= 1 // 除以2操作   -----------     i >>= 2 除以2操作 + 除以2操作
	fmt.Println(i)
}
