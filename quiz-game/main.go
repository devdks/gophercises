package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const usage = `
Usage of quiz-game:
	-f, --filename 	(string) A csv file in the format of 'question,answer' (default "problems.csv")
	-l, --limit 	(int)	 The time limit for the quiz in seconds (default is 5 seconds)
	-s ,--shuffle 	(bool) 	 Shuffle the questions
`

func main() {
	var score, limit int
	var csvFileName string
	var shuffle bool
	timer := time.NewTimer(time.Duration(limit) * time.Second)
	timer.Stop()
	ansCh := make(chan string)

	// Set up flags
	flag.IntVar(&limit, "limit", 5, "The time limit for the quiz in seconds (default is 5 seconds)")
	flag.IntVar(&limit, "l", 5, "The time limit for the quiz in seconds (default is 5 seconds)")
	flag.StringVar(&csvFileName, "filename", "problems.csv", "A csv file in the format of 'question,answer'")
	flag.StringVar(&csvFileName, "f", "problems.csv", "A csv file in the format of 'question,answer'")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle the questions")
	flag.BoolVar(&shuffle, "s", false, "Shuffle the questions")
	flag.Usage = func() {
		fmt.Print(usage)
	}
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
	records, err := csv.ReadAll()
	if err != nil {
		panic(err)
	}

	if shuffle {
		// Shuffle the records
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	// Loop through the records
	for idx, record := range records {

		question := record[0]
		answer := strings.TrimSpace(record[1])

		fmt.Println("--------------------")
		fmt.Printf("Problem #%d: %s = ", idx+1, question)

		go func() {
			timer.Reset(time.Duration(limit) * time.Second)

			var ans string
			ans, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				panic(err)
			}
			ansCh <- strings.TrimSpace(ans)
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up")
			continue
		case ans := <-ansCh:
			// fmt.Printf("%s, %s\n", ans, answer)
			if strings.EqualFold(ans, answer) {
				score++
			}
		}
		fmt.Println("--------------------")

	}
	fmt.Printf("✨✨ You have scored %d out of %d. ✨✨", score, len(records))
}
