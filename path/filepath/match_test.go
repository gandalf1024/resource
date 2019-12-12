package filepath

import (
	"fmt"
	"path/filepath"
	"testing"
)

//filepath包实现了兼容各操作系统的文件路径的实用操作函数。

func Test_Match(t *testing.T) {
	mat, err := filepath.Match("*", "cccccc")
	if err != nil {
		panic(err)
	}
	fmt.Println(mat)
}

func Test_Glob(t *testing.T) {
	ms, err := filepath.Glob("/99")
	if err != nil {
		panic(err)
	}
	fmt.Println(ms)
}

func Test_HasPrefix(t *testing.T) {
	flag := filepath.HasPrefix("11sadfsad", "11s")
	fmt.Println(flag)
}

func Test_Phone(t *testing.T) {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var count = 0
	for index := 0; index < len(num); index++ {
		i := index
		for num[i+1] == num[i] {
			count++
			i++
		}

		if count <= 3 {
			count = 0
		}
	}

	if count >= 3 {
		fmt.Println("符合连续3个")
	}

}
