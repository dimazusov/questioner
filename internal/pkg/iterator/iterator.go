package iterator

import (
	"errors"
	"optimization/internal/pkg/sentence"
)

type QuestionIterator struct {
	sent      sentence.Sentence
	questions []sentence.Question
	index     int
}

func FindQuestions(s sentence.Sentence) []sentence.Question {
	var questions []sentence.Question
	question := new(sentence.Question)
	for _, w := range s.Words {
		if w.Word == "{" {
			question = new(sentence.Question)
		} else if w.Word == "}" {
			questions = append(questions, *question)
		} else {
			question.Words = append(question.Words, w)
		}
	}
	return questions
}

func NewQuestionIterator(s sentence.Sentence) *QuestionIterator {
	return &QuestionIterator{sent: s, questions: FindQuestions(s)}
}

func (q *QuestionIterator) Has() bool {
	return q.questions != nil
}

func (q *QuestionIterator) GetNextQuestion() sentence.Question {
	r := q.questions[0]
	if len(q.questions) == 1 {
		q.questions = nil
	} else {
		q.questions = q.questions[1:]
	}
	return r
}

func (q *QuestionIterator) ReplaceFirstQuestion(response sentence.Sentence) error {
	result := new(sentence.Sentence)
	from, to := 0, 0
	for i, w := range q.sent.Words {
		if w.Word == "{" {
			from = i
		} else if w.Word == "}" {
			to = i + 1
			result.Words = append(result.Words, q.sent.Words[:from]...)
			result.Words = append(result.Words, response.Words...)
			result.Words = append(result.Words, q.sent.Words[to:]...)
			q.sent = *result
			return nil
		}
	}
	return errors.New("question { " + sentence.Sentence(q.questions[0]).Sentence() + " } \n\t was not replaced by the response { " + response.Sentence() + " }")
}

func (q *QuestionIterator) Sentence() *sentence.Sentence {
	q.sent.CountWord = uint(len(q.sent.Words))
	return &q.sent
}
