package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// open file and read contents
func readFile() [][]string {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("unable to open file")
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("unable to read file")
	}
	return records
}

func giveQuiz(quizFile [][]string){
	totalScore := 0

	for i, question := range quizFile {
		fmt.Printf("Question #%d: %s Enter your answer --> ",i + 1, question[0])

		// get user input
		var userAnswer int

		for {
			_, err := fmt.Scanf("%d", &userAnswer)
				if err == nil {
					break;
				}
				fmt.Println("not valid answer, try again")
				var dump string
				fmt.Scanln(&dump)
			
		}
		
		// check if answer is correct
		intAnswer, _ := strconv.Atoi(question[1])
		if userAnswer == intAnswer{
			totalScore++
		} 
		
	}
	fmt.Printf("You got a total of %d correct answers", totalScore)
}

func main() {
	quiz := readFile()
	giveQuiz(quiz)
	
}