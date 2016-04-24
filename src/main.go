package main

import (
	"fmt"
	"zhangmingkai4315/calc"
)

func main() {
	var x, y int = 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
}
