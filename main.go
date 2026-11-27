package main

import (
	"io"
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"bytes"
	"flag"
)

// dont overcomplicate it, map from q -> a, if user_a is incorrect, log it
// plan:
// take in conditional cli args, print the result of inputs
// load csv into program
// print first line of csv and take input

// func getQuizFromArgs() string {
// 	var quizName = flag.String("fileName", "./problems.csv", "Enter the filename of the quiz you'd like to attempt")
// 	flag.Parse()
// 	return quizName
// }

func buildMapFromCSV(fileName string) map[string]string {
	in, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Couldn't read file: " + fileName, err)
	}

	r := csv.NewReader(bytes.NewReader(in))

	m := make(map[string]string)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		m[record[0]] = record[1]
		// fmt.Println(record)
		// fmt.Printf("%T\n", record)
	}

	return m
}

func getFlags() string {
	quizName := flag.String("fileName", "./problems.csv", "Enter the full filepath of the quiz you'd like to attempt")
	flag.Parse()
	return *quizName
}

func main() {
	quizName := getFlags()
	fmt.Println(quizName)
	questionMap := buildMapFromCSV(quizName)

	for k, v := range questionMap {
		fmt.Println(k, v)
	}
}
