package main

import (
	"fmt"
	"github.com/divin3circle/qeldlie/internal/examples"
	"math/big"
)

func main() {
	// Create field elements in GF(17)
	a, _ := examples.NewFieldElement(3, 17)
	b, _ := examples.NewFieldElement(5, 17)

	// Addition: 3 + 5 = 8 (mod 17)
	sum, _ := a.Add(b)
	fmt.Printf("Sum: %s\n", sum)

	// Multiplication: 3 * 5 = 15 (mod 17)
	prod, _ := a.Mul(b)
	fmt.Printf("Product: %s\n", prod)

	// Exponentiation: 3^5 = 243 = 5 (mod 17)
	pow, _ := a.Pow(5)
	fmt.Printf("Power: %s\n", pow)

	// Division: 3/5 = 3 * 5^(17-2) = 3 * 5^15 (mod 17)
	div, _ := a.Div(b)
	fmt.Printf("Division: %s\n", div)

	// Negative exponent: 3^(-2) in GF(17)
	negPow, _ := a.Pow(-2)
	fmt.Printf("Negative power: %s\n", negPow)

	// Large field example
	largePrime, _ := new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	largeNum, _ := new(big.Int).SetString("37246545362547456373635472", 10)

	largeField, _ := examples.NewFieldElementFromBigInt(largeNum, largePrime)
	largeSquared, _ := largeField.Pow(2)
	fmt.Printf("Large field calculation: %s\n", largeSquared)
}
