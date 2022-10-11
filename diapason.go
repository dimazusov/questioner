package main

import "optimization/internal/pkg/sentence"

type Diapason struct {
	Sent       sentence.Question
	IsQuestion bool
}

func extractQuestions(s sentence.Sentence) []Diapason {
	var diapasons []Diapason
	question := new(sentence.Question)

	for _, word := range s.Words {
		switch word.Word {
		case "{":
			question.CountWord = uint(len(question.Words))
			if question.Words != nil {
				diapasons = append(diapasons, Diapason{Sent: *question})
			}
			question = &sentence.Question{ID: s.ID}
		case "}":
			question.CountWord = uint(len(question.Words))
			diapasons = append(diapasons, Diapason{
				Sent:       *question,
				IsQuestion: true,
			})
			question = &sentence.Question{ID: s.ID}
		default:
			question.Words = append(question.Words, word)
		}
	}

	return diapasons
}
