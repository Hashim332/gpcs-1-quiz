package main

import (
	"io"
	"strings"
	"fmt"
	"encoding/csv"
	"os"
	"log"
)

// dont overcomplicate it, map from q -> a, if user_a is incorrect, log it
// plan:
// take in conditional cli args, print the result of inputs
// load csv into program
// print first line of csv and take input

func main() {
	// data, err := os.ReadFile("./problems.csv")
	
	in, err := os.ReadFile("./problems.csv")
	if err != nil {
		log.Fatal(err)
	}


	r := csv.NewReader(strings.NewReader(string(in)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
		fmt.Println(len(record))
		// fmt.Println(record[0])
		// fmt.Println(record[1])
	}

	// fmt.Printf("%T\n", data)
	// fmt.Println(string(data))
}
