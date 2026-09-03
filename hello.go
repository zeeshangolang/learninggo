package main

import (
	"fmt"
)

type FuncOpts struct {
	Name string
	age  int
}

func FFUnc(funcopts FuncOpts) {
	fmt.Print(funcopts.Name, funcopts.age)
}

func main() {

	FFUnc(FuncOpts{
		Name: "zeeshan",
		age:  21,
	})

}
