package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Type the filename or use the default problems.csv!")

	var filename string
	fmt.Scanf("%s", &filename)
	if filename == "" {
		filename = "problems.csv"
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open the file", filename)
	}

	fmt.Println("Press a key when you are ready...")
	var key string
	fmt.Scanf("%s", &key)

	done := make(chan bool)

	timer := time.NewTimer(5 * time.Second)

	var correctAnswers int
	var failedAnswers int
	go func() {
		quiz(file, &correctAnswers, &failedAnswers)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("\nCompleted the quiz")
		break
	case <-timer.C:
		fmt.Println("\nTime has run out")
		break
	}

	fmt.Println("You failed ", failedAnswers, " questions")
	fmt.Println("You guesses ", correctAnswers, " questions")

}

func quiz(file *os.File, correctAnswers *int, failedAnswers *int) {
	reader := csv.NewReader(file)

	var answer string

	for {
		record, err := reader.Read()
		if err != nil {
			log.Println("File fully read.")
			break
		}

		fmt.Print("How much is ", record[0], "? \nA: ")

		fmt.Scanf("%s", &answer)

		if answer != record[1] {
			*failedAnswers++
			fmt.Println("Failed")
		} else {
			fmt.Println("Correct")
			*correctAnswers++
		}
	}
}
