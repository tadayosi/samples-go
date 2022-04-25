package types

import (
	"fmt"
	"testing"
)

// https://qiita.com/sxarp/items/cd528a546d1537105b9d

type A struct {
	Value bool
}

func (a A) Match(c Case) { c.A(a) }

type B struct {
	Value int
}

func (b B) Match(c Case) { c.B(b) }

type Case struct {
	A func(A)
	B func(B)
}

// Union type
type AB interface {
	Match(Case)
}

// Function using an union type
func print(ab AB) {
	ab.Match(Case{
		A: func(a A) { fmt.Printf("A: %v\n", a.Value) },
		B: func(b B) { fmt.Printf("B: %v\n", b.Value) },
	})
}

func TestUnion(t *testing.T) {
	a := A{Value: true}
	b := B{Value: 123}
	print(a)
	print(b)
}
