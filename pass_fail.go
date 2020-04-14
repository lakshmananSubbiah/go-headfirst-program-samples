//pass fail reports whether a grade is passing or failing
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
	fmt.Print("enter a value")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	parsedinput := strings.TrimSpace(input)
	numberInput,err := strconv.ParseFloat(parsedinput, 64)
	if err != nil{
		log.Fatal(err)
	}
	var status string
	if numberInput >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}

	fmt.Println("A grade of ",numberInput," is ",status)
}
