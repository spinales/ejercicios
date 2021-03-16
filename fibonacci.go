package main

func fibonacci(value int) int {
	a, b := 0, 1
	for i := 1; i < value; i++ {
		b = a + b
		a = b - a
	}
	return a
}
