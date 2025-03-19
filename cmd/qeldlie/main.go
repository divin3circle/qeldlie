package main

import (
	"fmt"
	"github.com/divin3circle/qeldlie/internal/examples"
)

func main() {
	a, err := examples.NewFieldElement(9, 19)

	if err != nil {
		fmt.Errorf("error %v", err)
	}
	b, err := a.Pow(12)
	if err != nil {
		fmt.Errorf("error %v", err)
	}
	fmt.Println(b.String())
}
