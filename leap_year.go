package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a year")
	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	value := strings.TrimSpace(input)
	year, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	if year%4 == 0{
		if year%100 == 0 {
			if year%400 == 0{
				fmt.Println("Leap Year")
			} else {
				fmt.Println(" Not a Leap Year")
			}
		} else {
			fmt.Println("Leap Year")
		}
	} else {
		fmt.Println("Not a Leap Year")
	}


}
