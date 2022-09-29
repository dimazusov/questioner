package main

import (
	"context"
	"optimization/internal/pkg/sentence"
)

type Repository interface {
	GetResponseTemplate(ctx context.Context, q sentence.Template) (*sentence.Template, error)
	GetResponse(ctx context.Context, r sentence.Template) (*sentence.Sentence, error)
}

type MorphClient interface { // ???
}

type questioner struct {
	rep    Repository
	client MorphClient
}

func NewQuestionerAction(rep Repository, client MorphClient) *questioner {
	return &questioner{rep, client}
}

func (m questioner) Handle(ctx context.Context, s *sentence.Sentence) (judgments *sentence.Sentence, err error) {
	var (
		questions = extractQuestions(*s)
		sIter     = s.IntoIter()
		response  = sentence.Sentence{ID: s.ID}
		_         = response
		responses []sentence.Sentence
	)
	for _, q := range questions { // 1
		questionTemplate := sentence.Template{
			Sentence: sentence.Sentence(q),
			Left:     true,
		}
		responseTemplate, err := m.rep.GetResponseTemplate(ctx, questionTemplate)
		if err != nil {
			return nil, err
		}
		response, err := m.rep.GetResponse(ctx, *responseTemplate)
		if err != nil {
			return nil, err
		}
		responses = append(responses, *response)
	}
	for sIter.HasNext() {
		word := sIter.GetNext()
		_ = word
	}
	return nil, nil
}

func extractQuestions(s sentence.Sentence) []sentence.Question {
	var (
		sIter     = s.IntoIter()
		questions []sentence.Question
		question  sentence.Question
	)
	for sIter.HasNext() {
		word := sIter.GetNext()
		if word.Word == "{" {
			question = sentence.Question{ID: s.ID}
		} else if word.Word == "}" {
			question.CountWord = uint(len(question.Words))
			questions = append(questions, question)
		} else {
			question.Words = append(question.Words, word)
		}
	}
	return questions
}

func changeCase(form sentence.Form, wordCase string) sentence.Form { // неверно
	form.Word = form.NormalForm
	form.Tag.Case = wordCase
	return form
}
