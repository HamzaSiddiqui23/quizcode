package main

import (
	"fmt"

	"github.com/HamzaSiddiqui23/pkg"
)

func main() {
	var a, b int = 1, 2
	fmt.Println("1+2=", pkg.Add(a, b))
	fmt.Println("2-1=", pkg.Subtract(b, a))
	fmt.Println("1*2=", pkg.Multiply(a, b))
}
