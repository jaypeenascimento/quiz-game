package main

import (
	"os"

	quiz "github.com/jaypasnascimento/quiz-game/domain/quiz"
)

func main() {
	q, err := quiz.NewQuiz(
		[]quiz.Question{
			{Statement: "Cade o Iscas?", Answer: "Foi pra ponte..."},
			{Statement: "Cade o Breno?", Answer: "Foi pro pote"},
		})
	if err != nil {
		panic("Failed to start quiz")
	}

	err = q.WriteStatement(os.Stdout, 0)
	if err != nil {
		panic("Failed to write statement to stdout")
	}
}
