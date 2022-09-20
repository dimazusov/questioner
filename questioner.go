package main

import (
	"context"
	"fmt"
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
	for _, q := range questions {
		fmt.Println(q.Sentence.Sentence())
	}
	_ = questions
	return nil, nil
}

func deepCopy(sent sentence.Sentence) sentence.Sentence {
	newSentence := sentence.Sentence{
		ID:        sent.ID,
		CountWord: sent.CountWord,
	}
	newWords := make([]sentence.Form, len(sent.Words))
	for i, word := range sent.Words {
		tag := word.Tag
		newTag := sentence.Tag{
			POS:          checkAndCopy(tag.POS),
			Animacy:      checkAndCopy(tag.Animacy),
			Aspect:       checkAndCopy(tag.Aspect),
			Case:         checkAndCopy(tag.Case),
			Gender:       checkAndCopy(tag.Gender),
			Involvment:   checkAndCopy(tag.Involvment),
			Mood:         checkAndCopy(tag.Mood),
			Number:       checkAndCopy(tag.Number),
			Person:       checkAndCopy(tag.Person),
			Tense:        checkAndCopy(tag.Tense),
			Transitivity: checkAndCopy(tag.Transitivity),
			Voice:        checkAndCopy(tag.Voice),
		}
		newForm := sentence.Form{
			ID:                 word.ID,
			JudgmentID:         word.JudgmentID,
			Word:               word.Word,
			NormalForm:         word.NormalForm,
			Score:              word.Score,
			PositionInSentence: word.PositionInSentence,
			Tag:                newTag,
		}
		newWords[i] = newForm
	}
	newSentence.Words = newWords
	return newSentence
}

func checkAndCopy(str *string) *string {
	if str != nil {
		s := *str
		return &s
	}
	return nil
}

func extractQuestions(sent sentence.Sentence) []*sentence.Question {
	s := deepCopy(sent)
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
