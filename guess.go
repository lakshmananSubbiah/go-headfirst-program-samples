//guess challenges players to guess a random number
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("I have guessed a number between 1 and 100")
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Can you guess it")

	var success bool = false	
	for i := 0; i < 10; i++ {
		fmt.Print("Make a Guess: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		convvalue, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if convvalue > target {
			fmt.Println("Oops. Your Guess was High")
		} else if convvalue < target {
			fmt.Println("Oops. Your Guess was Low")
		} else {
			success = true
			fmt.Println("Good Job! You Guessed it right!")
			break
		}
	}
	if(!success){
		fmt.Println(" Sorry you missed it Exact Target Value: ", target)
	}
}
