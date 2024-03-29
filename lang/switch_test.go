package lang

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSwitch1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	c := rand.Intn(3)
	v := 0
	switch c {
	case 0:
		v++
	case 1:
		v++
	case 2:
		v++
	default:
		v = v + 100
	}
	assert.Equal(t, v, 1)
}

func TestSwitch2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	c := rand.Intn(3)
	v := 0
	switch c {
	case 0:
		v++
		fallthrough
	case 1:
		v++
		fallthrough
	case 2:
		v++
		fallthrough
	default:
		v = v + 100
	}
	fmt.Printf("c = %v, v = %v\n", c, v)
	assert.Greater(t, v, 100)
}

func TestSwitch3(t *testing.T) {
	a, b, c := 1, 2, 3
	switch {
	case a == 1 && b == 2:
		fmt.Println("a = 1, b = 2")
		// ok
	case a == 1:
		fmt.Println("a = 1")
		assert.Fail(t, "a = 1")
	case c == 3:
		fmt.Println("c = 3")
		assert.Fail(t, "c = 3")
	default:
		fmt.Println("default")
		assert.Fail(t, "default")
	}
}
