package main

import "fmt"

func main() {

	var numbers = make([]int, 3, 5)
	numbers = append(numbers, 1)
	fmt.Printf("len =%d , cap = %d , numbers = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("len =%d , cap = %d , numbers = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 3)
	fmt.Printf("len =%d , cap = %d , numbers = %v\n", len(numbers), cap(numbers), numbers)
	numbers[0] = 100
	s2 := make([]int, 3)
	copy(s2, numbers)

	fmt.Printf("len =%d , cap = %d , s2 = %v\n", len(s2), cap(s2), s2)

}
