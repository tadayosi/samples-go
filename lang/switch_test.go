package lang

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSwitch(t *testing.T) {
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
