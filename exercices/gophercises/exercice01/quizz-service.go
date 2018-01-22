package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type Quizz struct {
	timeLimit    time.Duration
	problems     []Problem
	currentScore int32
}

type score struct {
	// TODO use an enum
	stype    string
	duration int
	score    int
}

type Problem struct {
	q string
	a string
}


func (quizz *Quizz) Launch() {
	stopC := make(chan score, 1)

	go terminateQuizz(stopC, len(quizz.problems))
	go timer(stopC, quizz)

	perform(stopC, quizz)
}

func perform(stopC chan<- score, quizz *Quizz) {
	beginTime := time.Now()
	for i, p := range quizz.problems {
		if ask(p, i+1) {
			atomic.AddInt32(&quizz.currentScore, 1)
		}
	}

	currentScore := atomic.LoadInt32(&quizz.currentScore)
	endTime := time.Now()
	duration := endTime.Sub(beginTime)
	stopC <- score{stype: "Finished", duration: int(duration.Seconds()), score: int(currentScore)}

	// TODO: enhance this
	time.Sleep(3 * time.Hour)
}

func timer(stopC chan<- score, quizz *Quizz) {
	time.Sleep(quizz.timeLimit)
	currentScore := int(atomic.LoadInt32(&quizz.currentScore))
	stopC <- score{stype: "TimeOut", duration: int(quizz.timeLimit), score: currentScore}
}

func terminateQuizz(stopC <-chan score, maxScore int) {

	currentScore := <-stopC

	var msg string
	if currentScore.stype == "Finished" {
		msg = fmt.Sprintf("Congratulation, you finished the quizz in %d seconds.\n", currentScore.score)
	} else {
		msg = fmt.Sprintf("Game Over, you have run out of time...")
	}

	fmt.Printf(msgTemplate, msg, currentScore.score, maxScore)
	os.Exit(1)
}

func ask(p Problem, nb int) bool {
	fmt.Printf("Question #%d: %s = ?\n", nb, p.q)

	var answer string
	fmt.Scanf("%s\n", &answer)

	if answer == p.a {
		fmt.Println("Correct")
		return true
	} else {
		fmt.Println("Wrong answer...")
		return false
	}
}

const (
	msgTemplate string = `**************************

%s

Your final score is %d/%d

**************************
`	
)

