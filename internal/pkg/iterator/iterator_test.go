package iterator

import (
	"github.com/stretchr/testify/require"
	"optimization/internal/pkg/sentence"
	"reflect"
	"testing"
)

func TestFindQuestion(t *testing.T) {
	sent := getSentence()
	expected := []sentence.Question{
		{Words: []sentence.Form{{Word: "c"}}},
		{Words: []sentence.Form{{Word: "e"}, {Word: "f"}}},
	}
	questions := FindQuestions(sent)
	require.Equal(t, true, reflect.DeepEqual(questions, expected))
}

func TestReplaceFirstQuestion(t *testing.T) {
	sent := getSentence()
	questions := NewQuestionIterator(sent)
	for questions.Has() {
		q := questions.GetNextQuestion()
		r := getResponse(q)
		err := questions.ReplaceFirstQuestion(r)
		require.Nil(t, err)
	}
	expected := getResultSentence()
	result := *questions.Sentence()
	require.Equal(t, true, reflect.DeepEqual(result, expected))
}

func getSentence() sentence.Sentence {
	return sentence.Sentence{Words: []sentence.Form{
		{Word: "a"},
		{Word: "b"},
		{Word: "{"},
		{Word: "c"},
		{Word: "}"},
		{Word: "d"},
		{Word: "{"},
		{Word: "e"},
		{Word: "f"},
		{Word: "}"},
	}}
}

func getResultSentence() sentence.Sentence {
	return sentence.Sentence{Words: []sentence.Form{
		{Word: "a"},
		{Word: "b"},
		{Word: "1"},
		{Word: "d"},
		{Word: "2"},
	}, CountWord: 5}
}

func getResponse(q sentence.Question) sentence.Sentence {
	m := make(map[string]sentence.Sentence)
	m["c "] = sentence.Sentence{Words: []sentence.Form{{Word: "1"}}}
	m["e f "] = sentence.Sentence{Words: []sentence.Form{{Word: "2"}}}
	return m[sentence.Sentence(q).Sentence()]
}
