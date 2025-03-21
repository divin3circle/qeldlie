package examples

import (
	"errors"
	"fmt"
	"math/big"
)

// FieldElement represents an element in a finite field GF(p)
type FieldElement struct {
	Num   *big.Int
	Prime *big.Int
}

var ErrNotInRange = errors.New("num not in range")
var ErrNotSameField = errors.New("cannot operate on elements from different fields")
var ErrFailedToCreateSet = errors.New("failed to create set")
var ErrDivisionByZero = errors.New("division by zero")

// NewFieldElement creates a new field element in GF(p)
func NewFieldElement(num, prime int64) (*FieldElement, error) {
	n := big.NewInt(num)
	p := big.NewInt(prime)

	// Check if prime is actually prime
	if !p.ProbablyPrime(20) {
		return nil, errors.New("modulus must be prime")
	}

	// Check if num is in range [0, p-1]
	if n.Sign() < 0 || n.Cmp(p) >= 0 {
		return nil, ErrNotInRange
	}

	return &FieldElement{
		Num:   n,
		Prime: p,
	}, nil
}

// NewFieldElementFromBigInt creates a new field element from big.Int values
func NewFieldElementFromBigInt(num, prime *big.Int) (*FieldElement, error) {
	// Check if prime is actually prime
	if !prime.ProbablyPrime(20) {
		return nil, errors.New("modulus must be prime")
	}

	// Create a copy to avoid sharing references
	numCopy := new(big.Int).Set(num)
	primeCopy := new(big.Int).Set(prime)

	// Ensure num is in range [0, p-1] by taking modulo p
	numCopy.Mod(numCopy, primeCopy)

	return &FieldElement{
		Num:   numCopy,
		Prime: primeCopy,
	}, nil
}

// String returns a string representation of the field element
func (f *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%s(%s)", f.Prime.String(), f.Num.String())
}

// Equal checks if two field elements are equal
func (f *FieldElement) Equal(other *FieldElement) bool {
	if other == nil {
		return false
	}
	return f.Num.Cmp(other.Num) == 0 && f.Prime.Cmp(other.Prime) == 0
}

// Add computes (f + other) mod prime
func (f *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if f.Prime.Cmp(other.Prime) != 0 {
		return nil, ErrNotSameField
	}

	// (a + b) mod p
	sum := new(big.Int).Add(f.Num, other.Num)
	sum.Mod(sum, f.Prime)

	return &FieldElement{
		Num:   sum,
		Prime: new(big.Int).Set(f.Prime),
	}, nil
}

// Sub computes (f - other) mod prime
func (f *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if f.Prime.Cmp(other.Prime) != 0 {
		return nil, ErrNotSameField
	}

	// (a - b) mod p
	// For negative results, we handle by adding p
	diff := new(big.Int).Sub(f.Num, other.Num)
	diff.Mod(diff, f.Prime)

	return &FieldElement{
		Num:   diff,
		Prime: new(big.Int).Set(f.Prime),
	}, nil
}

// Mul computes (f * other) mod prime
func (f *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if f.Prime.Cmp(other.Prime) != 0 {
		return nil, ErrNotSameField
	}

	// (a * b) mod p
	product := new(big.Int).Mul(f.Num, other.Num)
	product.Mod(product, f.Prime)

	return &FieldElement{
		Num:   product,
		Prime: new(big.Int).Set(f.Prime),
	}, nil
}

// Pow computes (f ^ exponent) mod prime
func (f *FieldElement) Pow(exponent int64) (*FieldElement, error) {
	// Convert exponent to big.Int
	exp := big.NewInt(exponent)

	// Handle negative exponents in finite fields using Fermat's Little Theorem
	// If n < 0, then x^n = x^(n mod (p-1)) in GF(p)
	if exponent < 0 {
		// Calculate p-1 (the order of the multiplicative group)
		pMinusOne := new(big.Int).Sub(f.Prime, big.NewInt(1))

		// Make exponent positive using modular arithmetic
		exp.Mod(exp, pMinusOne)
		exp.Add(exp, pMinusOne) // Ensure positive by adding p-1
		exp.Mod(exp, pMinusOne)
	}

	// Calculate f.Num^exp mod f.Prime
	result := new(big.Int).Exp(f.Num, exp, f.Prime)

	return &FieldElement{
		Num:   result,
		Prime: new(big.Int).Set(f.Prime),
	}, nil
}

// Div computes (f / other) mod prime
// In a finite field, this is f * other^(p-2) mod prime
func (f *FieldElement) Div(other *FieldElement) (*FieldElement, error) {
	if f.Prime.Cmp(other.Prime) != 0 {
		return nil, ErrNotSameField
	}

	// Check for division by zero
	if other.Num.Cmp(big.NewInt(0)) == 0 {
		return nil, ErrDivisionByZero
	}

	// In a finite field GF(p), division a/b = a * b^(p-2) mod p
	// Using Fermat's Little Theorem: b^(p-1) ≡ 1 (mod p)
	// So b^(p-2) is the multiplicative inverse of b

	// Calculate p-2
	pMinusTwo := new(big.Int).Sub(f.Prime, big.NewInt(2))

	// Calculate other^(p-2) mod prime (the multiplicative inverse)
	inverse := new(big.Int).Exp(other.Num, pMinusTwo, f.Prime)

	// Calculate f * inverse mod prime
	result := new(big.Int).Mul(f.Num, inverse)
	result.Mod(result, f.Prime)

	return &FieldElement{
		Num:   result,
		Prime: new(big.Int).Set(f.Prime),
	}, nil
}

// CreateSet creates a set of field elements in GF(p) by computing x^(p-1) for each x in [1,p-1]
func CreateSet(p int64) ([]*FieldElement, error) {
	var elements []*FieldElement

	prime := big.NewInt(p)
	if !prime.ProbablyPrime(20) {
		return nil, errors.New("modulus must be prime")
	}

	pMinusOne := new(big.Int).Sub(prime, big.NewInt(1))

	for i := int64(1); i < p; i++ {
		a, err := NewFieldElement(i, p)
		if err != nil {
			return nil, ErrFailedToCreateSet
		}

		// Compute a^(p-1)
		powResult := new(big.Int).Exp(a.Num, pMinusOne, prime)

		element := &FieldElement{
			Num:   powResult,
			Prime: new(big.Int).Set(prime),
		}

		elements = append(elements, element)
	}

	return elements, nil
}

// PrintSet prints the elements of a set
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

// Additional utility functions:

// Inverse finds the multiplicative inverse of the field element
func (f *FieldElement) Inverse() (*FieldElement, error) {
	// Check if the element is zero
	if f.Num.Cmp(big.NewInt(0)) == 0 {
		return nil, ErrDivisionByZero
	}

	// Calculate p-2
	pMinusTwo := new(big.Int).Sub(f.Prime, big.NewInt(2))

	// Calculate f^(p-2) mod p
	inverse := new(big.Int).Exp(f.Num, pMinusTwo, f.Prime)

	return &FieldElement{
		Num:   inverse,
		Prime: new(big.Int).Set(f.Prime),
	}, nil
}

// AddIdentity returns the additive identity (zero) for this field
func (f *FieldElement) AddIdentity() *FieldElement {
	return &FieldElement{
		Num:   big.NewInt(0),
		Prime: new(big.Int).Set(f.Prime),
	}
}

// MulIdentity returns the multiplicative identity (one) for this field
func (f *FieldElement) MulIdentity() *FieldElement {
	return &FieldElement{
		Num:   big.NewInt(1),
		Prime: new(big.Int).Set(f.Prime),
	}
}
