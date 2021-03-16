package main

import "fmt"

func primos(value int) {
	if value < 1 {
		fmt.Println("No califica")
	} else if checkPrimo(value) {
		fmt.Println("es primo")
	} else {
		fmt.Println("es compuesto")
	}
}

func checkPrimo(value int) bool {
	for i := 0; i < value; i++ {
		if i < value && i > 1 && value%i == 0 {
			return false
		}
	}
	return true
}
