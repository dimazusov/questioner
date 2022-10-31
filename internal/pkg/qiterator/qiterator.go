package qiterator

import (
	"optimization/internal/pkg/sentence"
)

type QuestionIterator struct {
	questions []sentence.Question
	index     int
}

func NewQuestionIterator(s sentence.Sentence) *QuestionIterator {
	return &QuestionIterator{questions: sentence.FindQuestions(s)}
}

func (q *QuestionIterator) Has() bool {
	return q.index <= len(q.questions)
}

func (q *QuestionIterator) GetNextQuestion() sentence.Question {
	r := q.questions[0]
	if len(q.questions) == 1 {
		q.questions = nil
	} else {
		q.questions = q.questions[1:]
	}
	q.index++
	return r
}
