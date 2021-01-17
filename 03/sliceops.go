package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v ,lend=%d , cap=%d\n", s, len(s), cap(s))
}
func main() {
	var s []int // Zero value for slice is nil
	fmt.Println(s)
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	printSlice(s3)

	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice ")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(front)
	fmt.Println(tail)

}
