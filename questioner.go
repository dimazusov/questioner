package main

import (
	"context"
	"optimization/internal/pkg/sentence"
)

type Repository interface {
	GetResponseTemplate(ctx context.Context, q sentence.Template) (*sentence.Template, error)
	GetResponse(ctx context.Context, r sentence.Template) (*sentence.Sentence, error)
}

type questioner struct {
	rep Repository
}

func NewQuestionerAction(rep Repository) *questioner {
	return &questioner{rep}
}

func (m questioner) Handle(ctx context.Context, s *sentence.Sentence) (response *sentence.Sentence, err error) {
	var (
		questions = extractQuestions(*s)
		responses []sentence.Sentence
		words     []sentence.Form
	)
	response = &sentence.Sentence{ID: s.ID}
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
	var ( // дурно пахнущий кусок
		responseCount int
		isQuestion    bool
	)
	for _, word := range s.Words {
		switch word.Word {
		case "{":
			words = append(words, responses[responseCount].Words...)
			isQuestion = true
			responseCount++
		case "}":
			isQuestion = false
		default:
			if isQuestion {
				continue
			} else {
				words = append(words, word)
			}
		}
	}
	response.Words = words
	response.CountWord = uint(len(words))
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
