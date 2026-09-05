package main

import (
	"fmt"
	"strconv"
)

var opmap = map[string]func(int, int) int{

	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func main() {

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Print("Invaild expression", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print("error on p1")
			continue
		}
		operation := expression[1]
		opFunc, ok := opmap[operation]
		if !ok {
			fmt.Print("unsupported operation")
			continue
		}

		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(err)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)

	}

}

func add(i, j int) int { return i + j }

func subtract(i, j int) int { return i - j }

func multiply(i, j int) int { return i * j }

func divide(i, j int) int { return i / j }
