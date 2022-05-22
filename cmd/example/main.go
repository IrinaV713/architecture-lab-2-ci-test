package main

import (
	"flag"
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

func main() {
	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}
