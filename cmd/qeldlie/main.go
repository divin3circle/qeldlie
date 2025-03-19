package main

import (
	"fmt"
	"github.com/divin3circle/qeldlie/internal/examples"
	"os"
)

func main() {
	field, err := examples.NewFieldElement(17, 13)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "NewFieldElement: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", field.String())

}
