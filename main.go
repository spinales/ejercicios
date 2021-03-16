package main

import (
	"fmt"
)

func main() {
	var input int

	fmt.Println("Numero: ")
	fmt.Scanln(&input)

	for i := 1; i < input; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(value int) string {
	if value%15 == 0 {
		return "FIZZ BUZZ"
	} else if value%3 == 0 {
		return "FIZZ"
	} else if value%5 == 0 {
		return "BUZZ"
	}
	return ""
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
