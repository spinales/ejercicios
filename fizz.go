package main

import "fmt"

func fizzBuzz(value int) string {
	if value%15 == 0 {
		return "FIZZ BUZZ"
	} else if value%3 == 0 {
		return "FIZZ"
	} else if value%5 == 0 {
		return "BUZZ"
	}
	return ""
}

func fizzBuzzSecuencia(value int) {
	for i := 1; i < value; i++ {
		fmt.Println(fizzBuzz(i))
	}
}
