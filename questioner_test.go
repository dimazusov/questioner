package main

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"log"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"optimization/internal/pkg/sentence"
)

func TestNewConcepterAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fullSentence := getSentence("необходимо выполнить mv {какое имя у файла или папки который нужно переместить?} {какое имя у файла или папки в которую нужно переместить?}")
	expectedSentence := getSentence("необходимо выполнить mv 1.txt folder")

	rep := NewMockRepository(ctrl)
	client := NewMockMorphClient(ctrl)
	rep.EXPECT().
		GetQuestionTemplate(context.Background(), fullSentence).
		Times(1).
		Return(nil, nil)
	rep.EXPECT().
		GetResponseTemplate(context.Background(), nil).
		Times(1).
		Return(nil, nil)
	rep.EXPECT().
		GetResponse(context.Background(), nil).
		Times(1).
		Return(nil, nil)
	c := NewQuestionerAction(rep, client)
	givenSentence, err := c.Handle(context.Background(), &fullSentence.Sentence)
	require.Nil(t, err)
	require.Equal(t, true, reflect.DeepEqual(givenSentence, []sentence.Sentence{expectedSentence.Sentence}))
}

func getSentence(str string) sentence.Template {
	m := make(map[string]string)
	// fullStr
	m["необходимо выполнить mv {какое имя у файла или папки который нужно переместить?} {какое имя у файла или папки в которую нужно переместить?}"] = `{
	"left": true,
	"right": false,
	"sentence": {
		"id": 0,
		"count_words": 26,
		"words": [{
			"word": "необходимо",
			"normalForm": "необходимо",
			"score": 0.5,
			"positionInSentence": 0,
			"tag": {
				"pos": "PRED",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "pres",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "выполнить",
			"normalForm": "выполнить",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "INFN",
				"animacy": "",
				"aspect": "perf",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "tran",
				"voice": ""
			}
		}, {
			"word": "mv",
			"normalForm": "mv",
			"score": 0.75,
			"positionInSentence": 0,
			"tag": {
				"pos": "",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "{",
			"normalForm": "{",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "какое",
			"normalForm": "какой",
			"score": 0.6,
			"positionInSentence": 0,
			"tag": {
				"pos": "ADJF",
				"animacy": "",
				"aspect": "",
				"case": "accs",
				"gender": "neut",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "имя",
			"normalForm": "имя",
			"score": 0.6,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "accs",
				"gender": "neut",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "у",
			"normalForm": "у",
			"score": 0.9959,
			"positionInSentence": 0,
			"tag": {
				"pos": "PREP",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "файла",
			"normalForm": "файл",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "gent",
				"gender": "masc",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "или",
			"normalForm": "или",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "CONJ",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "папки",
			"normalForm": "папка",
			"score": 0.25,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "gent",
				"gender": "femn",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "который",
			"normalForm": "который",
			"score": 0.171232,
			"positionInSentence": 0,
			"tag": {
				"pos": "ADJF",
				"animacy": "inan",
				"aspect": "",
				"case": "accs",
				"gender": "masc",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "нужно",
			"normalForm": "нужно",
			"score": 0.666666,
			"positionInSentence": 0,
			"tag": {
				"pos": "PRED",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "pres",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "переместить",
			"normalForm": "переместить",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "INFN",
				"animacy": "",
				"aspect": "perf",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "tran",
				"voice": ""
			}
		}, {
			"word": "}",
			"normalForm": "}",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "{",
			"normalForm": "{",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "какое",
			"normalForm": "какой",
			"score": 0.6,
			"positionInSentence": 0,
			"tag": {
				"pos": "ADJF",
				"animacy": "",
				"aspect": "",
				"case": "accs",
				"gender": "neut",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "имя",
			"normalForm": "имя",
			"score": 0.6,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "accs",
				"gender": "neut",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "у",
			"normalForm": "у",
			"score": 0.9959,
			"positionInSentence": 0,
			"tag": {
				"pos": "PREP",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "файла",
			"normalForm": "файл",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "gent",
				"gender": "masc",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "или",
			"normalForm": "или",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "CONJ",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "папки",
			"normalForm": "папка",
			"score": 0.25,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "gent",
				"gender": "femn",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "в",
			"normalForm": "в",
			"score": 0.999327,
			"positionInSentence": 0,
			"tag": {
				"pos": "PREP",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "которую",
			"normalForm": "который",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "ADJF",
				"animacy": "",
				"aspect": "",
				"case": "accs",
				"gender": "femn",
				"involvement": "",
				"mood": "",
				"number": "sing",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "нужно",
			"normalForm": "нужно",
			"score": 0.666666,
			"positionInSentence": 0,
			"tag": {
				"pos": "PRED",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "pres",
				"transitivity": "",
				"voice": ""
			}
		}, {
			"word": "переместить",
			"normalForm": "переместить",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "INFN",
				"animacy": "",
				"aspect": "perf",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "tran",
				"voice": ""
			}
		}, {
			"word": "}",
			"normalForm": "}",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "",
				"animacy": "",
				"aspect": "",
				"case": "",
				"gender": "",
				"involvement": "",
				"mood": "",
				"number": "",
				"person": "",
				"tense": "",
				"transitivity": "",
				"voice": ""
			}
		}]
	}
}`
	// expectedStr
	m["необходимо выполнить mv 1.txt folder"] = `{
	"left": true,
	"right": false,
	"sentence": {
		"id": 0,
		"count_words": 5,
		"words": [{
				"word": "необходимо",
				"normalForm": "необходимо",
				"score": 0.5,
				"positionInSentence": 0,
				"tag": {
					"pos": "PRED",
					"animacy": "",
					"aspect": "",
					"case": "",
					"gender": "",
					"involvement": "",
					"mood": "",
					"number": "",
					"person": "",
					"tense": "pres",
					"transitivity": "",
					"voice": ""
				}
			},
			{
				"word": "выполнить",
				"normalForm": "выполнить",
				"score": 1.0,
				"positionInSentence": 0,
				"tag": {
					"pos": "INFN",
					"animacy": "",
					"aspect": "perf",
					"case": "",
					"gender": "",
					"involvement": "",
					"mood": "",
					"number": "",
					"person": "",
					"tense": "",
					"transitivity": "tran",
					"voice": ""
				}
			},
			{
				"word": "mv",
				"normalForm": "mv",
				"score": 0.75,
				"positionInSentence": 0,
				"tag": {
					"pos": "",
					"animacy": "",
					"aspect": "",
					"case": "",
					"gender": "",
					"involvement": "",
					"mood": "",
					"number": "",
					"person": "",
					"tense": "",
					"transitivity": "",
					"voice": ""
				}
			},
			{
				"word": "1.txt",
				"normalForm": "1.txt",
				"score": 1.0,
				"positionInSentence": 0,
				"tag": {
					"pos": "",
					"animacy": "",
					"aspect": "",
					"case": "",
					"gender": "",
					"involvement": "",
					"mood": "",
					"number": "",
					"person": "",
					"tense": "",
					"transitivity": "",
					"voice": ""
				}
			},
			{
				"word": "folder",
				"normalForm": "folder",
				"score": 1.0,
				"positionInSentence": 0,
				"tag": {
					"pos": "",
					"animacy": "",
					"aspect": "",
					"case": "",
					"gender": "",
					"involvement": "",
					"mood": "",
					"number": "",
					"person": "",
					"tense": "",
					"transitivity": "",
					"voice": ""
				}
			}
		]
	}
}`

	t := sentence.Template{}
	err := json.Unmarshal([]byte(m[str]), &t)
	if err != nil {
		log.Fatalln(err)
	}
	return t
}
