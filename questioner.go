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
	/// res := перемести {что?} {куда?}
	// resSent := перемести {что?} {куда?}
	//
	//for resSent.HasQuestions() {
	//	question := resSent.FirstQuestion()
	//	response := getResponse(question)
	//	resSent = resSent.ReplaceFirstQuestion(response)
	//}

	res := s.Copy().Iterator()
	for res.HasQuestion() {
		question := res.GetNextQuestion()
		responseTemplate, err := m.rep.GetResponseTemplate(ctx, question)
		if err != nil {
			return nil, err
		}
		response, err := m.rep.GetResponse(ctx, *responseTemplate)
		if err != nil {
			return nil, err
		}
		err = res.ReplaceFirstQuestion(*response)
		if err != nil {
			return nil, err
		}
	}
	return res.Sentence(), nil
}
