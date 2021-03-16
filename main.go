package main

import (
	"flag"
	"fmt"
	"sort"
)

func main() {
	ordena := flag.String("ordena", "", "ordena una palabra o una oracion.")
	flag.Parse()
	fmt.Println(orderingByte(*ordena))

	revr := flag.String("reverse", "", "imprime una palabra al reves.")
	flag.Parse()
	fmt.Println(reverse(*revr))

	primo := flag.Int("primo", -1, "determina si un numero es primo o no.")
	flag.Parse()
	if *primo >= 0 {
		fmt.Println(FizzBuzz(*primo))
	}

	fizz := flag.Int("fizzbuzz", 0, `si el numero actual es divisible entre 3, imprimir FIZZ. si el numero actual es divisible entre 5, imprimir BUZZ
	 si el numero actual de divisible entre 15, imprimir FIZZ BUZZ`)
	flag.Parse()
	if *fizz > 0 {
		fmt.Println(FizzBuzz(*fizz))
	}

	fibo := flag.Int("fibonacci", 0, "secuencia fibonacci, imprime el enesimo numero del numero dado.")
	flag.Parse()
	if *fibo > 0 {
		fmt.Println(Fibonacci(*fibo))
	}
}

func orderingByte(value string) string {
	arr := []byte(value)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}

func reverse(value string) string {
	result := ""
	for i := len(value); i > 0; i-- {
		result = result + string(value[i-1])
	}
	return result
}

func Fibonacci(value int) int {
	a, b, c := 0, 1, 0
	for i := 1; i < value; i++ {
		c = a + b
		a = b
		b = c
	}
	return a
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
