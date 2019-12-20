package list

import (
	"container/list"
	"fmt"
	"testing"
)

type user struct {
	name string
	age  int
}

func Test_list(t *testing.T) {
	l := list.New()
	l.PushBack(user{name: "zj_1", age: 11})
	l.PushBack(user{name: "zj_2", age: 22})
	l.PushBack(user{name: "zj_3", age: 33})
	l.PushBack(user{name: "zj_4", age: 44})
	l.PushBack(user{name: "zj_5", age: 55})
	l.PushFront(user{name: "zj_6", age: 66})

	l.InsertBefore(user{name: "zj_7", age: 77}, l.Back())

	for e := l.Front(); e != nil; e = e.Next() {
		le := e.Value.(user)
		fmt.Println(le)
	}

}
