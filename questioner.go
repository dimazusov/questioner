package main

import (
	"context"
	"optimization/internal/pkg/iterator"
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
	questions := iterator.NewQuestionIterator(s.Copy())
	for questions.Has() {
		question := questions.GetNextQuestion()
		responseTemplate, err := m.rep.GetResponseTemplate(ctx, question)
		if err != nil {
			return nil, err
		}
		response, err := m.rep.GetResponse(ctx, *responseTemplate)
		if err != nil {
			return nil, err
		}
		if err = questions.ReplaceFirstQuestion(*response); err != nil {
			return nil, err
		}
	}
	return questions.Sentence(), nil
}
