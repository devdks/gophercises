package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	var score, count, limit int
	var csvFileName string

	// Set up flags
	flag.IntVar(&limit, "limit", -1, "The time limit for the quiz in seconds (default is 10 seconds)")
	flag.StringVar(&csvFileName, " filename", "problems.csv", "A csv file in the format of 'question,answer'")
	flag.StringVar(&csvFileName, "f", "problems.csv", "A csv file in the format of 'question,answer'")
	// Parse the flags
	flag.Parse()

	// Open the csv file
	file, err := os.Open(csvFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new csv reader
	csv := csv.NewReader(file)

	// Loop through the csv file
	for {
		// Read the csv file
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		count++

		var problem Problem
		problem.Question = record[0]
		problem.Answer = strings.TrimSpace(record[1])

		fmt.Printf("Problem #%d: %s = ", count, problem.Question)
		var answer string
		fmt.Scanln(&answer)
		if answer == problem.Answer {
			score++
		}

	}
	fmt.Printf("You have scored %d out of %d.", score, count)
}
