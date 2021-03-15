package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Println("Numero: ")
	fmt.Scanln(&input)

	if input < 1 {
		fmt.Println("No califica")
	} else if primo(input) {
		fmt.Println("es primo")
	} else {
		fmt.Println("es compuesto")
	}
}

func primo(value int) bool {
	for i := 0; i < value; i++ {
		if i < value && i > 1 && value%i == 0 {
			return false
		}
	}
	return true
}
