package ch09

import "fmt"

func Sub() {
	list_9_1()
}

func list_9_1() {
	s1 := store()                                // s1 is closure
	s2 := store()                                // s2 is closure
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // (x, n) = (0, 1)(=1), (0 ,2)(=2)
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // (x, n) = (1, 1)(=2), (2 ,2)(=4)
	fmt.Printf("s1: %d, s2: %d\n", s1(1), s2(2)) // (x, n) = (2, 1)(=3), (4 ,2)(=6)

}

func store() func(int) int {
	var x int
	return func(n int) int {
		x += n
		return x
	}
}
