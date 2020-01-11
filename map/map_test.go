package _map

import (
	"fmt"
	"testing"
)

func Test_map(t *testing.T) {
	var m map[string]string
	m = make(map[string]string, 0)
	m["11"] = "aa"

	var s []int
	s = append(s, 11)
	fmt.Println(s)
}
