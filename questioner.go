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

func (m questioner) Handle(ctx context.Context, s *sentence.Sentence) (res *sentence.Sentence, err error) {
	diapasons := extractQuestions(*s)
	res = &sentence.Sentence{ID: s.ID}
	for _, d := range diapasons {
		if d.IsQuestion {
			responseTemplate, err := m.rep.GetResponseTemplate(ctx, d.Sent)
			if err != nil {
				return nil, err
			}
			response, err := m.rep.GetResponse(ctx, *responseTemplate)
			if err != nil {
				return nil, err
			}
			res.Words = append(res.Words, response.Words...)
		} else {
			res.Words = append(res.Words, d.Sent.Words...)
		}
	}
	res.CountWord = uint(len(res.Words))
	return res, nil
}
