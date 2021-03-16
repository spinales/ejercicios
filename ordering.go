package main

import "sort"

func orderingByte(value string) string {
	arr := []byte(value)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return string(arr)
}
