package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	DefaultCsvFileName = "problems.csv"
	DefaultTimeLimit = 30
)


func main() {
	
	// Define known flags
	csvFileName := flag.String("f", DefaultCsvFileName, "a csv file name where question and answer are defined, one per line separated with a coma")
	timeLimitInS := flag.Int("t", DefaultTimeLimit, "the time limit in seconds")
	
	// timeLimitInSec := flag.Duration("t", DefaultTimeLimit, "the time limit in seconds")
	
	// Parse the flags
	flag.Parse()
	
	// Retrieve Q&A from the CSV file 
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll();
	if err != nil {
		exit("Canot parse the provided CSV file")
	}
	
	// We now have a slice of question and answers
	currProblems := parseLines(lines)
	
	maxDuration := time.Duration(*timeLimitInS * 1000000 * 1000) 
	
	// fmt.Printf("Input time limit:  %d ->  %v\n", *timeLimitInS, maxDuration)

	quizz := Quizz{timeLimit: maxDuration, problems: currProblems}
	quizz.Launch()
}

func parseLines(lines [][]string) ([]Problem) {
	fmt.Printf("Preparing the quizz with %d questions\n", len(lines))
	var problems = make ([]Problem, len(lines))
	for i, line := range lines {
		// fmt.Printf("Line %d: question = %s, answer = %s\n", i, line[0], line[1])
		problems[i] = Problem{q: line[0], a: strings.TrimSpace(line[1])}
	}	
	return problems
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}
