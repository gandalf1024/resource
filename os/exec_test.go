package os

import (
	"fmt"
	"os"
	"testing"
)

func Test_exec(t *testing.T) {
	// 启动一个wait 程序   ps aux | grep main 获取pid
	pro, err := os.FindProcess(49042)
	if err != nil {
		panic(err)
	}
	err = pro.Kill()
	if err != nil {
		panic(err)
	}

	fmt.Println(pro.Pid)
}

func Test_StartProcess(t *testing.T) {

	f, err := os.Open("/home/zhangjun/golang/src/resource/temp/d_01/main")
	if err != nil {
		panic(err)
	}

	files := make([]*os.File, 0)
	files = append(files, f)

	pro := &os.ProcAttr{Dir: "/home/zhangjun/golang/src/resource/temp/d_01/", Files: files}
	process, err := os.StartProcess("", nil, pro)
	if err != nil {
		panic(err)
	}
	fmt.Println(process)

}
