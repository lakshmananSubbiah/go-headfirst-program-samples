package main

import (
	"fmt"
)

func main() {
	fmt.Println(1.0 / 3.0)
	fmt.Printf(" %f \n", 1.0/3.0)
	fmt.Printf(" %12f \n", 1.0/3.0)
	value := fmt.Sprintf(" %f \n", 1.0/3.0)
	fmt.Printf("%12.3f \n",65.0/21.0)
	fmt.Println(value)
}
