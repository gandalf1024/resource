package flag

import (
	"flag"
	"fmt"
	"testing"
)

func Test_flag(t *testing.T) {
	var ip = flag.Int("flagname", 1234, "help message for flagname")
	fmt.Println(ip)
}
