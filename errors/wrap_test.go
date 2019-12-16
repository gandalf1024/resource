package errors

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Unwrap(t *testing.T) {
	err := errors.New("123456")

	err = errors.Unwrap(err)
	fmt.Println(err)

}
