package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func stats(count, total, correct, wrong int) {
	wrong = total - (correct + wrong)
	fmt.Printf("\n--- End of quiz ---\n")
	fmt.Printf("Total answered: %d / %d questions\n", count, total)
	fmt.Println("Correct: ", correct)
	fmt.Println("Incorrect: ", wrong)
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "file to set the questions from")
	timeInSeconds := flag.Int("timer", 30, "set the timer for the quiz in seconds")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hit enter to start the quiz")
	text, _ := reader.ReadString('\n')
	if text != "\n" {
		fmt.Println("You've chosen to exit")
		os.Exit(0)
	}

	total := 0
	count := 0
	correct := 0
	wrong := 0

	timer := time.AfterFunc(time.Duration(*timeInSeconds)*time.Second, func() {
		fmt.Printf("\nQuiz time (%d seconds) has elapsed\n", *timeInSeconds)
		stats(count, total, correct, wrong)
		os.Exit(0)
	})
	defer timer.Stop()

	file, err := os.Open(*csvFile)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)

	fmt.Printf("\n--- Start of quiz ---\n")
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	total = len(records)
	for _, record := range records {
		count++
		fmt.Printf("%d. %s\n", count, record[0])
		fmt.Print("Enter your answer: ")
		text, _ := reader.ReadString('\n')
		if text == record[1]+"\n" {
			correct++
		} else {
			wrong++
		}
	}
	stats(count, total, correct, wrong)
}
