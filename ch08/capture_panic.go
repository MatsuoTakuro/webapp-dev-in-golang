package ch08

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func list_8_14() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("caputured panic: %v\n", r)
		}
	}()

	fmt.Println("output")
	panic("happening")
	fmt.Println("not output")
}

func list_8_15() {
	ns, err := Parse("1 2 3 4 5")
	if err != nil {
		log.Fatalln(err)
	}

	for _, n := range ns {
		fmt.Print(n, " ")
	}
}

func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	panic("123")
	return
}

func fields2numbers(fs []string) []int {
	var numbers []int
	for _, f := range fs {
		n, _ := strconv.ParseInt(f, 10, 32)
		numbers = append(numbers, int(n))
	}
	return numbers
}
