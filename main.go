package main

import (
	"fmt"
	// "encoding/csv"
	"os"
	"log"
)

// dont overcomplicate it, map from q -> a, if user_a is incorrect, log it
// plan:
// take in conditional cli args, print the result of inputs
// load csv into program
// print first line of csv and take input

func main() {
	data, err := os.ReadFile("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	
	// os.Stdout.Write(data)
	fmt.Println("%T/n", string(data))
}
