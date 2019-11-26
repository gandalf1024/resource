package os

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func Test_Notify(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)
}

func Test_Stop(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Stop(c)

	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)

}
