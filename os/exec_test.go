package os

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func Test_exec(t *testing.T) {
	// 启动一个wait 程序   ps aux | grep main 获取pid
	pro, err := os.FindProcess(49042) // 循环调用在linux下会出现问题但是不会死机
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

func Test_Command(t *testing.T) {
	cmd := exec.Command("shutdown", "-h", "now") // 对应不同平台命令
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func Test_String(t *testing.T) {
	cmd := exec.Command("shutdown", "-h", "now") // 对应不同平台命令
	fmt.Println(cmd.String())
}

func Test_LookPath(t *testing.T) {
	path, err := exec.LookPath("./data/data.txt")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)
}
