package main

import (
	"encoding/csv"
	"log"
	"os"
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

func main() {
	quiz := readFile()
	
}