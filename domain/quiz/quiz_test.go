package quiz_test

import (
	"bytes"
	"testing"

	"github.com/jaypasnascimento/quiz-game/domain/quiz"
)

func TestQuizWriteStatement(t *testing.T) {
	q, err := quiz.NewQuiz([]quiz.Question{
		{Statement: "Cade o Iscas?", Answer: "Foi pra ponte..."},
		{Statement: "Cade o Breno?", Answer: "Foi pro pote"},
	})
	if err != nil {
		t.Fatal("Failed to initialize Quiz")
	}

	t.Run("it write statement", func(t *testing.T) {
		want := "Cade o Iscas?"
		output := bytes.Buffer{}

		q.WriteStatement(&output, 0)

		if want != output.String() {
			t.Errorf("want %q, got %q", want, output.String())
		}
	})

	t.Run("it write statement for question idx 1", func(t *testing.T) {
		want := "Cade o Breno?"
		output := bytes.Buffer{}

		q.WriteStatement(&output, 1)

		if want != output.String() {
			t.Errorf("want %q, got %q", want, output.String())
		}
	})

	t.Run("it return error when idx is not found", func(t *testing.T) {
		err := q.WriteStatement(&bytes.Buffer{}, 999)

		want := quiz.NotFoundErr

		if want != err {
			t.Errorf("want %v, but got %v", want, err)
		}
	})
}
