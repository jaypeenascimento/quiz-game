// Package quiz: contains the logic for the quiz game.
package quiz

import (
	"errors"
	"fmt"
	"io"
)

var NotFoundErr = errors.New("Question not found")

type Question struct {
	Statement string
	Answer    string
}

type Quiz struct {
	questions []Question
}

func (q Quiz) WriteStatement(w io.Writer, idx int) error {
	statement := q.getQuestion(idx)
	if statement == ""{
		return NotFoundErr
	}

	_, err := fmt.Fprintf(w, "%s", statement)
	return err
}

func (q Quiz) getQuestion(idx int) string {
	if idx >= len(q.questions) {
		return ""
	}

	return q.questions[idx].Statement
}

func NewQuiz(questions []Question) (Quiz, error) {
	return Quiz{
		questions: questions,
	}, nil
}
