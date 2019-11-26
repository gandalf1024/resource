package os

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"sync"
	"syscall"
	"testing"
	"time"
)

func Test_Open(t *testing.T) {
	f, err := os.Open("./data/data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Name())

	bys, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bys))

}
func Test_Open_syscall(t *testing.T) {
	_, _ = os.Open("./data/data.txt")
	_, _ = os.Open("./data/data1.txt")
	_, _ = os.Open("./data/data2.txt")

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			bys := make([]byte, 1000)
			_, err := syscall.Read(3, bys)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(bys))
			wg.Done()
		}()
	}
	wg.Wait()

}

func Test_Read(t *testing.T) {
	f, err := os.Open("./data/data.txt")
	if err != nil {
		panic(err)
	}

	bys := make([]byte, 80)
	n, err := f.Read(bys[20:])
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(string(bys))
}

func Test_ReadAt(t *testing.T) {
	f, err := os.Open("./data/data.txt")
	if err != nil {
		panic(err)
	}
	bys := make([]byte, 100)

	n, err := f.ReadAt(bys, 10)
	if err != nil {
		if err != io.EOF { //读取完成err = EOF 注意不是报错
			panic(err)
		}
	}
	fmt.Println(n)
	fmt.Println(string(bys))
}

func Test_Write(t *testing.T) {
	f, err := os.OpenFile("./data/data.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	bys := []byte("123456789")
	//err =  f.Chmod(os.ModePerm)
	//if err != nil {
	//	panic(err)
	//}
	n, err := f.Write(bys)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}

func Test_WriteAt(t *testing.T) {
	f, err := os.OpenFile("./data/data.txt", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.WriteAt([]byte("个人缺乏而戈尔甘"), 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}

func Test_Seek(t *testing.T) {
	//f, err := os.OpenFile("./data/data.txt", os.O_WRONLY|os.O_APPEND, os.ModePerm)
	f, err := os.Open("./data/data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾。
	//whence为底层汇编参数
	ret, err := f.Seek(35, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)

	bys := make([]byte, 100)
	n, err := f.Read(bys)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "------------------------------")
	fmt.Println(string(bys))
}

func Test_WriteString(t *testing.T) {
	f, err := os.OpenFile("./data/data.txt", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	n, err := f.WriteString("457345dfhgsdfhsdfhdsf第三粉红色的规范节能司法德国\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func Test_Mkdir(t *testing.T) {
	err := os.Mkdir("./data2", os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func Test_Rename(t *testing.T) {
	err := os.Rename("./data3", "./data5")
	if err != nil {
		panic(err)
	}
}

func Test_TempDir(t *testing.T) {
	dir := os.TempDir()
	fmt.Println(dir)
}

func Test_UserCacheDir(t *testing.T) {
	str, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}

	fmt.Println(str)
}

func Test_UserConfigDir(t *testing.T) {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
}

func Test_UserHomeDir(t *testing.T) {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
}

func Test_Chmod(t *testing.T) {
	fmt.Println(os.ModeDir)
	fmt.Println(os.ModeDir)
	fmt.Println(os.ModeAppend)
	fmt.Println(os.ModeExclusive)
	fmt.Println(os.ModeTemporary)
	fmt.Println(os.ModeSymlink)
	fmt.Println(os.ModeDevice)
	fmt.Println(os.ModeNamedPipe)
	fmt.Println(os.ModeSocket)
	fmt.Println(os.ModeSetuid)
	fmt.Println(os.ModeSetgid)
	fmt.Println(os.ModeCharDevice)
	fmt.Println(os.ModeSticky)
	fmt.Println(os.ModeIrregular)
	fmt.Println(os.ModeType)
	fmt.Println(os.ModePerm)

	err := os.Chmod("./data/data.txt", os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// 错误
func Test_SetDeadline(t *testing.T) {
	f, err := os.Open("/home/zhangjun/soft/goland-2019.2.5.tar.gz")
	if err != nil {
		panic(err)
	}

	err = f.SetDeadline(time.Time{})
	if err != nil {
		panic(err)
	}

	bys, _ := ioutil.ReadAll(f)
	fmt.Println(string(bys))
}

// 错误
func Test_SetReadDeadline(t *testing.T) {
	f, err := os.OpenFile("./data/data.txt", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = f.SetReadDeadline(time.Time{})
	if err != nil {
		panic(err)
	}

}

func Test_SyscallConn(t *testing.T) {
	f, err := os.Open("./data/data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rawc, err := f.SyscallConn()
	if err != nil {
		panic(err)
	}

	err = rawc.Read(func(fd uintptr) (done bool) {
		fmt.Println(fd)
		return true
	})

	by := runtime.ReadTrace()
	fmt.Println(string(by))

}

func Test_Hostname(t *testing.T) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
}

func Test_Environ(t *testing.T) {
	envs := os.Environ()
	for _, v := range envs {
		fmt.Println(v)
	}
}

func Test_Exit(t *testing.T) {
	defer func() {
		fmt.Println("123456")
	}()

	fmt.Println("22222")
	os.Exit(0) // 相当于关电源
}

//字符串替换
func Test_Expand(t *testing.T) {
	cc := os.Expand("123456${var}", func(s string) string {
		fmt.Println(s)
		return "963"
	})
	fmt.Println(cc, "---")
}

func Test_ExpandEnv(t *testing.T) {
	result := os.ExpandEnv("ccc${var}")
	fmt.Println(result)
}
