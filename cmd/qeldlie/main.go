package main

import (
	"fmt"
	"os"

	"github.com/divin3circle/qeldlie/internal/examples"
)

func main() {
	p1 := &examples.Point{}
	p1, err := p1.CreateEllipticPoint(-1, -1, 5, 7)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error creating elliptic point: %s\n", err)
		os.Exit(1)
	}
	p2 := &examples.Point{}
	p2, err = p2.CreateEllipticPoint(-1, -1, 5, 7)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error creating elliptic point: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(p1.String())
	fmt.Println(p2.String())
	fmt.Println(p1.AdditiveIdentity(p2))

}
