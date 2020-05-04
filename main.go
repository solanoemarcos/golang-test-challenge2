package main

import "github.com/solanoemarcos/fallabella-fif-test-2/parser"

func main() {
	input := []byte("11AB398765UJ1A05")
	parser.TlvParse(input)
}
