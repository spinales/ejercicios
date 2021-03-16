package main

import (
	"flag"
	"fmt"
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
		primos(*primo)
	}

	fizz := flag.Int("fizzbuzz", 0, `si el numero actual es divisible entre 3, imprimir FIZZ. si el numero actual es divisible entre 5, imprimir BUZZ
	 si el numero actual de divisible entre 15, imprimir FIZZ BUZZ`)
	flag.Parse()
	if *fizz > 0 {
		fmt.Println(fizzBuzz(*fizz))
	}

	fibo := flag.Int("fibonacci", 0, "secuencia fibonacci, imprime el enesimo numero del numero dado.")
	flag.Parse()
	if *fibo > 0 {
		fmt.Println(fibonacci(*fibo))
	}
}
