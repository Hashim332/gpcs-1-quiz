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
func readCSV(fileName) {
	in, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Couldn't read file: " + fileName, err)
	}

	r := csv.NewReader(strings.NewReader(string(in)))
	r.Comma = " "

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
	}
}

func main() {
	readCSV("./problems.csv")
}
