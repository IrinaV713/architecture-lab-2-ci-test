package main

import (
	"fmt"
	lab2 "github.com/roman-mazur/architecture-lab-2"
)

func main() {
	res, _ := lab2.PostfixToPrefix("2 2 +")
	fmt.Println(res)
}
