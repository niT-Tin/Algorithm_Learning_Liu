package main

import "fmt"

func LinearSearch(tars [8]int, target int) int {
	for i, v := range tars {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	var tars = [...]int{24, 18, 12, 9, 16, 66, 32, 4}
	fmt.Println(LinearSearch(tars, 16))
}
