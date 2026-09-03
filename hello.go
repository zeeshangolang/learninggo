package main

import (
	"fmt"
)

func main() {

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i := range evenVals {
		evenVals[i] = 1
	}
	fmt.Println(evenVals)
}
