package main

import (
	"context"
	"optimization/internal/pkg/sentence"
)

type Repository interface {
	GetResponseTemplate(ctx context.Context, q sentence.Question) (*sentence.Template, error)
	GetResponse(ctx context.Context, r sentence.Template) (*sentence.Sentence, error)
}

type questioner struct {
	rep Repository
}

func NewQuestionerAction(rep Repository) *questioner {
	return &questioner{rep}
}

func (m questioner) Handle(ctx context.Context, s *sentence.Sentence) (response *sentence.Sentence, err error) {
	questions := extractQuestions(*s)
	responses := make([]sentence.Sentence, len(questions))
	sIter := IntoIter(s)
	response = &sentence.Sentence{ID: s.ID}
	for i, q := range questions {
		responseTemplate, err := m.rep.GetResponseTemplate(ctx, q)
		if err != nil {
			return nil, err
		}
		response, err := m.rep.GetResponse(ctx, *responseTemplate)
		if err != nil {
			return nil, err
		}
		responses[i] = *response
	}
	for sIter.HasNext() {
		word, isQuestion := sIter.GetNext()
		if isQuestion {
			response.Words = append(response.Words, responses[sIter.ResponseIdx].Words...)
		} else {
			response.Words = append(response.Words, word)
		}
	}
	response.CountWord = uint(len(response.Words))
	return response, nil
}

func extractQuestions(s sentence.Sentence) []sentence.Question {
	var (
		questions []sentence.Question
		question  sentence.Question
	)
	for _, word := range s.Words {
		switch word.Word {
		case "{":
			question = sentence.Question{ID: s.ID}
		case "}":
			question.CountWord = uint(len(question.Words))
			questions = append(questions, question)
		default:
			question.Words = append(question.Words, word)
		}
	}
	return questions
}
