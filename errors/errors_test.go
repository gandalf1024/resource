package errors

import (
	"fmt"
	"testing"
)

type my_err struct {
	desc string
}

func (*my_err) Error() string {
	return "mkv"
}

func (*my_err) NIHAO() error {
	return &my_err{desc: "kkk"}
}

func Test_New(t *testing.T) {
	var me interface{} = &my_err{} //类型提升 向上转

	u, ok := me.(interface { // 判断是否实现NIHAO接口
		NIHAO() error
	})

	fmt.Println(u)
	fmt.Println(ok)
}
