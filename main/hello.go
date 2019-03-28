package main

import (
	"fmt"

	"github.com/fatih/color"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	color.Cyan("Prints text in cyan.")
}
