package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/avast/retry-go"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var count int

func flaky() error {
	count++
	fmt.Printf("count = %d\n", count)
	i := rand.Intn(3)
	if i > 0 {
		return fmt.Errorf("i = %d", i)
	}
	return nil
}

func TestRetry(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	count = 0

	err := retry.Do(flaky, retry.Attempts(10))
	assert.Nil(t, err)
}

func TestEventually(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	count = 0

	g := NewWithT(t)
	g.Eventually(flaky).Should(BeNil())
}
