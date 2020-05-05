package main

import (
	"fmt"

	"github.com/solanoemarcos/golang-test-challenge2/parser"
)

func main() {
	input := []byte("11AB398765UJ1A05")
	result, err := parser.TlvParse(input)
	if err != nil {
		fmt.Print("Unexpected error")
	}
	fmt.Printf("Lenght: %s \n", result["length"])
	fmt.Printf("Type: %s \n", result["type"])
	fmt.Printf("Value: %s \n", result["value"])
}
