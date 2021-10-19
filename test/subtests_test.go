package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSubtests(t *testing.T) {
	t.Run("A", func(t *testing.T) {
		t.Run("A1", func(t *testing.T) {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("Test A1")
		})
		t.Run("A2", func(t *testing.T) {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Test A2")
		})
		t.Run("A3", func(t *testing.T) {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Test A3")
		})
		time.Sleep(300 * time.Millisecond)
		fmt.Println("Test A")
	})
	t.Run("B", func(t *testing.T) {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Test B")
	})
	t.Run("C", func(t *testing.T) {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Test C")
	})
}
