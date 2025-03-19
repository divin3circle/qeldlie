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
var ErrNotSameField = errors.New("cannot add nums from different fields")
var ErrFailedToCreateSet = errors.New("failed to create set")

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
func (f *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if other.Prime != f.Prime {
		return nil, ErrNotSameField
	}

	num := (f.Num + other.Num) % f.Prime
	return NewFieldElement(num, f.Prime)
}

func (f *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if other.Prime != f.Prime {
		return nil, ErrNotSameField
	}

	if f.Num < other.Num {
		// a -f b = (a - b) mod p
		// -n mod k = K - (n mod K)
		num := f.Prime - (other.Num-f.Num)%f.Prime
		return NewFieldElement(num, f.Prime)
	}
	num := (f.Num - other.Num) % f.Prime
	return NewFieldElement(num, f.Prime)
}

func (f *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if f.Prime != other.Prime {
		return nil, ErrNotSameField
	}
	num := (f.Num * other.Num) % f.Prime
	return NewFieldElement(num, f.Prime)
}

func (f *FieldElement) Pow(exp int) (*FieldElement, error) {
	result := 1
	for i := 0; i < exp; i++ {
		result = (result * f.Num) % f.Prime
	}

	return NewFieldElement(result, f.Prime)
}

func CreateSet(p int) ([]*FieldElement, error) {
	var result []*FieldElement

	for i := 1; i < p; i++ {
		a, err := NewFieldElement(i, p)

		if err != nil {
			return nil, ErrFailedToCreateSet
		}
		b, err := a.Pow(p - 1)
		if err != nil {
			return nil, err
		}
		result = append(result, b)
	}
	return result, nil
}

func PrintSet(set []*FieldElement) {
	fmt.Print("Set: [")
	for idx, elem := range set {
		if idx == len(set)-1 {
			fmt.Print(elem.String(), "]")
			return
		}
		fmt.Print(elem.String(), ", ")
	}
	fmt.Println("]")
}
