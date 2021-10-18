package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/avast/retry-go"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	count := 0
	err := retry.Do(func() error {
		count++
		fmt.Printf("count = %d\n", count)
		i := rand.Intn(3)
		if i > 0 {
			return fmt.Errorf("i = %d", i)
		}
		return nil
	})
	assert.Nil(t, err)
}
