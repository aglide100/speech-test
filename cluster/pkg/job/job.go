package job

import (
	"errors"

	"github.com/jdkato/prose/v2"
)


type Job struct {
	Content string
	Speaker string
	Id	string
	No int
}

func DivideTest(text string) ([]string, error) {
	doc, _ := prose.NewDocument(text)
	
	sents := doc.Sentences()
	
	texts := []string{}

	for _, sent := range sents {
		if len(sent.Text) > 255 {
			return nil, errors.New("Sent too long : /" + sent.Text)
		}
		texts = append(texts, sent.Text)
	}
	return texts, nil
}