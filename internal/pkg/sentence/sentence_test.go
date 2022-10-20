package sentence

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	sent := getSentence()
	clone := *sent.Copy()
	require.Equal(t, true, reflect.DeepEqual(sent, clone))
}

func TestSentence(t *testing.T) {
	sent := getSentence()
	str := sent.Sentence()
	expected := "a b { c } d { e f } "
	require.Equal(t, str, expected)
}

func TestFindQuestion(t *testing.T) {
	sent := getSentence()
	expected := []Question{
		{Words: []Form{{Word: "c"}}},
		{Words: []Form{{Word: "e"}, {Word: "f"}}},
	}
	questions := FindQuestions(sent)
	require.Equal(t, true, reflect.DeepEqual(questions, expected))
}

func getSentence() Sentence {
	return Sentence{Words: []Form{
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
