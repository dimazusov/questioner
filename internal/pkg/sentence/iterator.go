package sentence

import "errors"

type QuestionIterator struct {
	sentence  Sentence
	questions []Question
	index     int
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

func (s Sentence) Iterator() QuestionIterator {
	return QuestionIterator{sentence: s, questions: s.FindQuestions()}
}

func (q *QuestionIterator) HasQuestion() bool {
	return q.questions != nil
}

func (q *QuestionIterator) GetNextQuestion() Question {
	r := q.questions[0]
	if len(q.questions) == 1 {
		q.questions = nil
	} else {
		q.questions = q.questions[1:]
	}
	return r
}

func (q *QuestionIterator) ReplaceFirstQuestion(response Sentence) error {
	result := new(Sentence)
	from, to := 0, 0
	for i, w := range q.sentence.Words {
		if w.Word == "{" {
			from = i
		} else if w.Word == "}" {
			to = i + 1
			result.Words = append(result.Words, q.sentence.Words[:from]...)
			result.Words = append(result.Words, response.Words...)
			result.Words = append(result.Words, q.sentence.Words[to:]...)
			q.sentence = *result
			return nil
		}
	}
	return errors.New("question { " + Sentence(q.questions[0]).Sentence() + " } \n\t was not replaced by the response { " + response.Sentence() + " }")
}

func (q *QuestionIterator) Sentence() *Sentence {
	q.sentence.CountWord = uint(len(q.sentence.Words))
	return &q.sentence
}
