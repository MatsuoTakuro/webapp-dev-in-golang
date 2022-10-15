package ch09

import (
	"fmt"
	"sync"
)

func Sub() {
	// list_9_1()
	// list_9_3()
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

func list_9_3() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("i: %d\n", i)
			wg.Done()
		}()
	}
	wg.Wait()

	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(j int) {
			fmt.Printf("j: %d\n", j)
			wg.Done()
		}(j)
	}
	wg.Wait()

	for k := 0; k < 5; k++ {
		k := k
		wg.Add(1)
		go func() {
			fmt.Printf("k: %d\n", k)
			wg.Done()
		}()
	}
	wg.Wait()
	// i: 5
	// i: 5
	// i: 5
	// i: 5
	// i: 5
	// j: 4
	// j: 0
	// j: 1
	// j: 2
	// j: 3
	// k: 4
	// k: 0
	// k: 1
	// k: 2
}
