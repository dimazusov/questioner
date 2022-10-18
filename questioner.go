package main

import (
	"context"
	"optimization/internal/pkg/qiterator"
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

func (m questioner) Handle(ctx context.Context, s sentence.Sentence) (*sentence.Sentence, error) {
	questionIterator := qiterator.NewQuestionIterator(s.Copy())
	for questionIterator.Has() {
		question := questionIterator.GetNextQuestion()
		responseTemplate, err := m.rep.GetResponseTemplate(ctx, question)
		if err != nil {
			return nil, err
		}
		response, err := m.rep.GetResponse(ctx, *responseTemplate)
		if err != nil {
			return nil, err
		}
		newS, err := s.ReplaceFirstQuestion(*response)
		if err != nil {
			return nil, err
		}
		s = *newS
	}
	return questionIterator.Sentence(), nil
}
