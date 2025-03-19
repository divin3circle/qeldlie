package main

import (
	"fmt"
	"github.com/divin3circle/qeldlie/internal/examples"
)

func main() {
	result, err := examples.CreateSet(7)
	if err != nil {
		fmt.Errorf("%v\n", err)
	}

	examples.PrintSet(result)
}
