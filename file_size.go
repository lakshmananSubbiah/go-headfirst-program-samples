package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("/Users/lakshmanan-zt71/Documents/go/pass_fail.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
