//go:build aaa && bbb
// +build aaa,bbb

package tags

import (
	"fmt"
	"testing"
)

func TestAAA_BBB(t *testing.T) {
	fmt.Println("AAA BBB")
}
