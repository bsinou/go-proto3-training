package main

import (
	"fmt"
)

type Quizz struct {
	timeLimit int
	problems []Problem
	currentScore int 
}

type Problem struct {
	q string
	a string
}

func (quizz *Quizz) Perform(){
	for i, p := range quizz.problems {
		if ask(p, i+1) {
			quizz.currentScore ++
		}
	}
	fmt.Printf("**************************\nGame Over...\nYour final score is %d/%d\n**************************\n\n", quizz.currentScore, len(quizz.problems))
}

func ask(p Problem, nb int) (bool){
	fmt.Printf("Question #%d: %s = ?\n", nb, p.q)
	
	var answer string
	fmt.Scanf("%s\n",&answer)
	
	if answer == p.a {
		fmt.Println("Correct")
		return true
	} else {
		fmt.Println("Wrong answer...")
		return false
	}
}



