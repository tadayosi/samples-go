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

func flaky() error {
	i := rand.Intn(3)
	if i > 0 {
		return fmt.Errorf("i = %d", i)
	}
	return nil
}

func TestRetry(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	count := 0
	err := retry.Do(func() error {
		count++
		fmt.Printf("count = %d\n", count)
		return flaky()
	}, retry.Attempts(10))
	assert.Nil(t, err)
}

func TestGomegaRetry(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	g := NewGomegaWithT(t)

	count := 0
	err := retry.Do(func() error {
		count++
		fmt.Printf("count = %d\n", count)
		g.Expect(flaky).To(BeNil())
		return nil
	}, retry.Attempts(10))

	assert.Nil(t, err)
}
