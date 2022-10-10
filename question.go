package main

import "optimization/internal/pkg/sentence"

type QuestionIter struct {
	sent        sentence.Sentence
	wordIdx     int
	ResponseIdx int
}

func IntoIter(s *sentence.Sentence) *QuestionIter {
	return &QuestionIter{sent: *s, ResponseIdx: -1}
}

func (i *QuestionIter) HasNext() bool {
	return i.wordIdx < len(i.sent.Words)
}

func (i *QuestionIter) GetNext() (word sentence.Form, isQuestion bool) {
	word = i.sent.Words[i.wordIdx]
	if word.Word == "{" {
		for j, word := range i.sent.Words[i.wordIdx:] {
			if word.Word == "}" {
				i.wordIdx += j + 1
				i.ResponseIdx++
				return word, true
			}
		}
	}
	i.wordIdx++
	return word, false
}
