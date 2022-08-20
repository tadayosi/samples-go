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
	valueB string
}

func (b B) PrintB() {
	fmt.Printf("B: %s\n", b.valueB)
}

type C struct {
	A
	B      B
	valueC string
}

func (c C) PrintC() {
	fmt.Printf("C: %s, B: %s, A: %s\n", c.valueC, c.B.valueB, c.valueA)
}

func TestObject(t *testing.T) {
	c := C{A: A{valueA: "AAA"}, B: B{valueB: "BBB"}, valueC: "CCC"}
	c.PrintA()
	//c.PrintB() // compile error
	c.B.PrintB()
	c.PrintC()
}
