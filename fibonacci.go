package main

func fibonacci(value int) int {
	a, b, c := 0, 1, 0
	for i := 1; i < value; i++ {
		c = a + b
		a = b
		b = c
	}
	return a
}
