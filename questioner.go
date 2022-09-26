package main

import (
	"context"
	"optimization/internal/pkg/morph"
	"optimization/internal/pkg/sentence"
)

type Repository interface {
	GetQuestionTemplate(ctx context.Context, q sentence.Sentence) (*sentence.Template, error)
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
	questions := extractQuestions(*s)
	for i, q := range questions { // 1
		normalForm := sentence.Sentence{}
		for i, w := range q.Sentence.Words {
			if isNoun(w) {
				normalForm.Words = q.Sentence.Words[i:]
				normalForm.CountWord = q.Sentence.CountWord - uint(i)
				normalForm.ID = q.Sentence.ID
				break
			}
		}
		questions[i] = &sentence.Question{Sentence: normalForm}
	}
	return nil, nil
}

func extractQuestions(s sentence.Sentence) []*sentence.Question {
	var questions []*sentence.Question
	for i := 0; uint(i) < s.CountWord; i++ {
		if s.Words[i].Word == "{" {
			i++
			j := i
			for ; s.Words[j].Word != "}"; j++ {
			}
			words := make([]sentence.Form, j-i)
			for idx, word := range s.Words[i:j] { // фабричный метод?
				w := word
				w.Tag = word.Tag
				words[idx] = w
			}
			sent := sentence.Sentence{
				ID:        s.ID,
				CountWord: uint(len(words)),
				Words:     words,
			}
			question := sentence.Question{Sentence: sent}
			questions = append(questions, &question)
			i = j
		}
	}
	return questions
}

func isNoun(word sentence.Form) bool {
	if word.Tag.POS == "" {
		return false
	}
	return word.Tag.POS == morph.PartOfSpeachNOUN
}
