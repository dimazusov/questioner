package sentence

import "errors"

type Sentence struct {
	ID        uint   `json:"id" db:"id"`
	CountWord uint   `json:"count_words"`
	Words     []Form `json:"words" gorm:"foreignKey:JudgmentID"`
}

type Question Sentence

func (s Sentence) Sentence() string {
	var result string
	for _, word := range s.Words {
		result += word.Word + " "
	}
	return result
}

func (s Sentence) FindQuestions() []Question {
	var questions []Question
	question := new(Question)
	for _, w := range s.Words {
		if w.Word == "{" {
			question = new(Question)
		} else if w.Word == "}" {
			questions = append(questions, *question)
		} else {
			question.Words = append(question.Words, w)
		}
	}
	return questions
}

func (s Sentence) ReplaceQuestion(question Question, response Sentence) (Sentence, error) {
	result := new(Sentence)
	from, to := 0, 0
	for i, w := range s.Words {
		if w.Word == "{" {
			from = i
		} else if w.Word == "}" {
			to = i + 1
			result.Words = append(result.Words, s.Words[:from]...)
			result.Words = append(result.Words, response.Words...)
			result.Words = append(result.Words, s.Words[to:]...)
			return *result, nil
		}
	}
	return *result, errors.New("question { " + Sentence(question).Sentence() + " } \n\t was not replaced by the response { " + response.Sentence() + " }")
}

type Form struct {
	ID                 uint    `json:"id" db:"id"`
	JudgmentID         uint    `json:"judgmentID" db:"judgment_id"`
	Word               string  `json:"word" db:"word"`
	NormalForm         string  `json:"normalForm" db:"normal_form"`
	Score              float64 `json:"score" db:"score"`
	PositionInSentence int     `json:"positionInSentence" db:"position_in_sentence"`
	Tag                Tag     `json:"tag" db:"tag" gorm:"embedded;embeddedPrefix:tag_"`
}

type Tag struct {
	POS          string `json:"pos" db:"pos"`
	Animacy      string `json:"animacy" db:"animacy"`
	Aspect       string `json:"aspect" db:"aspect"`
	Case         string `json:"case" db:"case"`
	Gender       string `json:"gender" db:"gender"`
	Involvment   string `json:"involvment" db:"involvment"`
	Mood         string `json:"mood" db:"mood"`
	Number       string `json:"number" db:"number"`
	Person       string `json:"person" db:"person"`
	Tense        string `json:"tense" db:"tense"`
	Transitivity string `json:"transitivity" db:"transitivity"`
	Voice        string `json:"voice" db:"voice"`
}

type Indexes struct {
	I int
	J int
}

type Template struct {
	Sentence Sentence `json:"sentence" db:"sentence"`
	Left     bool     `json:"left" db:"left"`
	Right    bool     `json:"right" db:"right"`
}
