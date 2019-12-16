package log

import (
	"log"
	"os"
	"testing"
)

func Test_New(t *testing.T) {

	f, err := os.OpenFile("./temp/log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(f, "前缀_", 2)

	println(logger.Flags())
	logger.Output(2, "./temp")
	logger.Panic("1010")
	logger.Panicf("%d", 1111)
	logger.Panicln("1212")
	println(logger.Prefix())

	logger.Fatal("1234")

	logger.Fatalf("%d", 456)
	logger.Fatalln("789")
}
