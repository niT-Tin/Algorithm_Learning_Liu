package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var tars = [...]int{24, 18, 12, 9, 16, 66, 32, 4}
	fmt.Println(LinearSearchV1(tars, 16))
	var Stus = []Student{
		{"名字1", 20},
		{"名字2", 21},
		{"名字3", 22},
		{"名字4", 23},
		{"名字4", 22},
	}
	var s = Student{"名字1", 23}
	fmt.Println(LinearSearchV2(s, Stus))
}
