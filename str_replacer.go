package main

import (
	"fmt"
	"strings"
)

func main() {
	baseValue := "G# R#cks"
	replacer := strings.NewReplacer("#", "o")
	fmt.Println(replacer.Replace(baseValue))
}
