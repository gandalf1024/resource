package sort

import (
	"fmt"
	"sort"
	"testing"
)

func Test_Search(t *testing.T) {
	data := make([]int, 0, 10)
	data = append(data, 10)
	data = append(data, 20)
	data = append(data, 30)
	data = append(data, 15)
	data = append(data, 25)

	x := 23
	i := sort.Search(len(data), func(i int) bool {
		return data[i] >= x
	})

	fmt.Println(i)
}

func Test_SearchInts(t *testing.T) {
	data := make([]int, 0, 10)
	data = append(data, 10)
	data = append(data, 20)
	data = append(data, 30)
	data = append(data, 15)
	data = append(data, 25)

	indx := sort.SearchInts(data, 20)
	fmt.Println(indx)
}

func Test_SearchFloat64s(t *testing.T) {
	data := make([]float64, 0, 10)
	data = append(data, 10)
	data = append(data, 20)
	data = append(data, 30)
	data = append(data, 15)
	data = append(data, 25)

	indx := sort.SearchFloat64s(data, 30.0)
	fmt.Println(indx)
}
