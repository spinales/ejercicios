package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Println("Numero: ")
	fmt.Scanln(&input)

	primos(input)
}

func CheckPrimo(value int) bool {
	for i := 0; i < value; i++ {
		if i < value && i > 1 && value%i == 0 {
			return false
		}
	}
	return true
}

func primos(value int) {
	if value < 1 {
		fmt.Println("No califica")
	} else if CheckPrimo(value) {
		fmt.Println("es primo")
	} else {
		fmt.Println("es compuesto")
	}
}
