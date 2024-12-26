package main

import (
	"fmt"
	"strings"

	"github.com/AndreHammel/aah_golang_calculadora/calculator"
)

func main() {
	var input string
	fmt.Print("Enter the operation (using this format: 2*2):")
	fmt.Scan(&input)

	op := strings.Split(input, "")

	result, err := calculator.GetResult(op)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

}
