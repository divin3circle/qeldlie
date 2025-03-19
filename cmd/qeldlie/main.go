package main

import (
	"fmt"
	"github.com/divin3circle/qeldlie/internal/examples"
)

func main() {
	a, err := examples.NewFieldElement(5, 13)
	if err != nil {
		panic(err)
	}
	b, err := examples.NewFieldElement(7, 13)
	if err != nil {
		panic(err)
	}
	c, err := examples.NewFieldElement(11, 13)

	d, err := a.Sub(b)

	if err != nil {
		fmt.Errorf("error adding %v to %v: %v", b, a, err)
	}

	fmt.Println(d.Equal(c))
	fmt.Println(d.String())
}
