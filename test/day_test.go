package test

import (
	"fmt"
	"testing"
)

func Test_01(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
	}
}
