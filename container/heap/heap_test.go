package heap

import (
	"fmt"
	"sort"
	"testing"
)

type user struct {
	name string
	age  int
}

type user_heap struct {
	list []user
}

func (h *user_heap) Less(i, j int) bool {
	return (h.list)[i].age < (h.list)[j].age
}

func (h *user_heap) Swap(i, j int) {
	(h.list)[i], (h.list)[j] = (h.list)[j], (h.list)[i]
}

func (h *user_heap) Len() int {
	return len(h.list)
}

func (h *user_heap) Pop() (v interface{}) {
	h.list, v = (h.list)[:h.Len()-1], (h.list)[h.Len()-1]
	return
}

func (h *user_heap) Push(v interface{}) {
	h.list = append(h.list, v.(user))
}

func Test_K(t *testing.T) {
	heap_obj := user_heap{list: make([]user, 0)}
	heap_obj.list = append(heap_obj.list, user{name: "zj", age: 1})
	heap_obj.list = append(heap_obj.list, user{name: "zj2", age: 2})
	heap_obj.list = append(heap_obj.list, user{name: "zj3", age: 3})
	heap_obj.list = append(heap_obj.list, user{name: "zj4", age: 4})
	heap_obj.list = append(heap_obj.list, user{name: "zj5", age: 5})
	heap_obj.list = append(heap_obj.list, user{name: "zj6", age: 6})

	sort.Sort(&heap_obj)

	uinter := heap_obj.Pop()
	u := uinter.(user)
	fmt.Println(u)
	heap_obj.Push(u)
	uinter = heap_obj.Pop()
	u = uinter.(user)
	fmt.Println(u)

}
