package main

import "fmt"

func main() {

	fmt.Println("Aaary")

	var list = [2]int{1, 2}

	fmt.Println(list)

	fmt.Println("slice ")

	var list1 = []int{1, 2, 3, 4}

	fmt.Println(list1)

	fmt.Println("maps")

	s := make(map[string]int)

	s["sandeep"] = 1

	val, ok := s["sandeep"]

	fmt.Println(val, ok)

}
