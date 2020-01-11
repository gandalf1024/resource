package atomic

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test_Load(t *testing.T) {
	var v atomic.Value
	v.Store(11)
	v.Store(12) // 覆盖11
	v.Store(13) // 覆盖12

	result := v.Load() // 获取唯一值
	fmt.Println(result)
	result2 := v.Load() // 获取唯一值
	fmt.Println(result2)
	result3 := v.Load() // 获取唯一值
	fmt.Println(result3)
}
