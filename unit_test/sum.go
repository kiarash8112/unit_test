package main

import "fmt"

func Ints(vi ...int) int {
	return ints(vi)
}
func ints(vi []int) int {
	if len(vi) == 0 {
		return 0
	}
	return ints(vi[1:]) + vi[0]
}

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Project Started!")
	fmt.Println(ints(numbers))
}
