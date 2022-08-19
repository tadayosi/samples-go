package object

import (
	"fmt"
	"testing"
)

type A struct {
	valueA string
}

func (a A) PrintA() {
	fmt.Printf("A: %s\n", a.valueA)
}

type B struct {
	A
	valueB string
}

func (b B) PrintB() {
	fmt.Printf("B: %s, A: %s\n", b.valueB, b.valueA)
}

func TestObject(t *testing.T) {
	b := B{A: A{valueA: "AAA"}, valueB: "BBB"}
	b.PrintA()
	b.PrintB()
}
