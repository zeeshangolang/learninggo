package main

import (
	"fmt"
)

var opmap = map[string]func(int, int) int{

	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func main() {

	jprint := func(j int) {
		fmt.Printf("printing j from inside %d \n", j)
	}

	for i := 0; i < 5; i++ {
		jprint(i)
	}

}

func add(i, j int) int { return i + j }

func subtract(i, j int) int { return i - j }

func multiply(i, j int) int { return i * j }

func divide(i, j int) int { return i / j }
