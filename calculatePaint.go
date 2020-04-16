package main

import (
	"fmt"
)

var metersPerLitre float64
func main() {
	metersPerLitre = 10.0
	var total,paintNeeded float64
	paintNeeded = calculatePaint(5.0, 20.0)
	total += paintNeeded
	paintNeeded = calculatePaint(2.9, 2.3)
	total += paintNeeded
	paintNeeded = calculatePaint(28.9, 23.4)
	total += paintNeeded
	fmt.Printf("The total paint needed is : %0.2f litres \n",total)
	
}

func calculatePaint(width float64, length float64) float64 {
	area := width * length
	return area/metersPerLitre
}
