package lang

import (
	"fmt"
	"testing"
)

func TestFor1(t *testing.T) {
Loop1:
	for _, v := range []int{1, 2, 3, 4, 5} {
		if v == 2 {
			fmt.Print("x")
			continue
		}
		if v == 4 {
			fmt.Print("x")
			continue Loop1
		}
		fmt.Print(v)
	}
	fmt.Println()
}
