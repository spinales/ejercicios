package main

func reverse(value string) string {
	result := ""
	for i := len(value); i > 0; i-- {
		result = result + string(value[i-1])
	}
	return result
}
