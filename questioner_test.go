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

func TestNewQuestionerAction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fullSentence := getSentence("необходимо выполнить mv {какое имя у файла или папки который нужно переместить?} {какое имя у файла или папки в которую нужно переместить?}")
	expectedSentence := getSentence("необходимо выполнить mv 1.txt в folder")
	responseTemplate1 := getSentence("файл который нужно переместить")
	responseTemplate2 := getSentence("папка в которую нужно переместить")
	responses := []sentence.Template{getSentence("1.txt"), getSentence("в folder")}
	questions := fullSentence.Sentence.FindQuestions()

	rep := NewMockRepository(ctrl)

	rep.EXPECT().
		GetResponseTemplate(context.Background(), questions[0]).
		Times(1).
		Return(&responseTemplate1, nil)
	rep.EXPECT().
		GetResponse(context.Background(), responseTemplate1).
		Times(1).
		Return(&responses[0].Sentence, nil)

	rep.EXPECT().
		GetResponseTemplate(context.Background(), questions[1]).
		Times(1).
		Return(&responseTemplate2, nil)
	rep.EXPECT().
		GetResponse(context.Background(), responseTemplate2).
		Times(1).
		Return(&responses[1].Sentence, nil)

	c := NewQuestionerAction(rep)
	givenSentence, err := c.Handle(context.Background(), fullSentence.Sentence)
	require.Nil(t, err)
	require.Equal(t, true, reflect.DeepEqual(givenSentence, &expectedSentence.Sentence))
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
	m["необходимо выполнить mv 1.txt в folder"] = `{
	"left": true,
	"right": false,
	"sentence": {
		"id": 0,
		"count_words": 6,
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
	// responseTemplate 1
	m["файл который нужно переместить"] = `{
	"left": true,
	"right": false,
	"sentence": {
		"id": 0,
		"count_words": 4,
		"words": [{
			"word": "файл",
			"normalForm": "файл",
			"score": 0.5,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "nomn",
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
			"word": "который",
			"normalForm": "который",
			"score": 0.828767,
			"positionInSentence": 0,
			"tag": {
				"pos": "ADJF",
				"animacy": "",
				"aspect": "",
				"case": "nomn",
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
		}]
	}
}`
	// responseTemplate 2
	m["папка в которую нужно переместить"] = `{
	"left": true,
	"right": false,
	"sentence": {
		"id": 0,
		"count_words": 5,
		"words": [{
			"word": "папка",
			"normalForm": "папка",
			"score": 1.0,
			"positionInSentence": 0,
			"tag": {
				"pos": "NOUN",
				"animacy": "inan",
				"aspect": "",
				"case": "nomn",
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
		}]
	}
}`
	// response 1
	m["1.txt"] = `{
	"left": true,
	"right": false,
	"sentence": {
		"id": 0,
		"count_words": 1,
		"words": [{
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
			}
		]
	}
}`
	// response 2
	m["в folder"] = `{
	"left": true,
	"right": false,
	"sentence": {
		"id": 0,
		"count_words": 2,
		"words": [{
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
