package main

import (
	"flag"
	"fmt"
)

func main() {
	ordena := flag.String("ordena", "", "ordena una palabra o una oración.")
	revr := flag.String("reverse", "", "imprime una palabra al revés.")
	primo := flag.Int("primo", -1, "determina si un número es primo o no.")
	fizz := flag.Int("fizzbuzz", 0, `si el número actual es divisible entre 3, imprimir FIZZ. si el número actual es divisible entre 5, imprimir BUZZ
	 si el número actual de divisible entre 15, imprimir FIZZ BUZZ`)
	fibo := flag.Int("fibonacci", 0, "secuencia fibonacci, imprime el enésimo número del numero dado.")
	servr := flag.Bool("server", false, "inicia un servidor básico.")

	flag.Parse()

	fmt.Println(orderingByte(*ordena))
	fmt.Println(reverse(*revr))

	if *primo >= 0 {
		primos(*primo)
	}

	if *fizz > 0 {
		fizzBuzzSecuencia(*fizz)
	}

	if *fibo > 0 {
		fmt.Println(fibonacci(*fibo))
	}

	if *servr {
		server()
	}
}
