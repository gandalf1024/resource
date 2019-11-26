package os

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

func Test_LookPath(t *testing.T) {
	path, err := exec.LookPath("./data/data.txt")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)
}
