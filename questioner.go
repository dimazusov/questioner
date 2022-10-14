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

func (m questioner) Handle(ctx context.Context, s sentence.Sentence) (result *sentence.Sentence, err error) {
	res := sentence.Sentence{Words: s.Words}
	questions := s.FindQuestions()
	for _, q := range questions {
		responseTemplate, err := m.rep.GetResponseTemplate(ctx, q)
		if err != nil {
			return nil, err
		}
		response, err := m.rep.GetResponse(ctx, *responseTemplate)
		if err != nil {
			return nil, err
		}
		newRes, err := res.ReplaceQuestion(q, *response)
		if err != nil {
			return nil, err
		}
		res = newRes
	}
	res.CountWord = uint(len(res.Words))
	return &res, nil
}
