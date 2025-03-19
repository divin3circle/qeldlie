package examples

import (
	"errors"
	"fmt"
)

type FieldElement struct {
	Num   int
	Prime int
}

var ErrNoInRange = errors.New("num not in range")

func NewFieldElement(num, prime int) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, ErrNoInRange
	}
	return &FieldElement{
		Num:   num,
		Prime: prime,
	}, nil
}

func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", f.Prime, f.Num)
}

func (f *FieldElement) Equal(other *FieldElement) bool {
	if other == nil {
		return false
	}
	return f.Num == other.Num && f.Prime == other.Prime
}
