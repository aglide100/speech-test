package job

import (
	"github.com/jdkato/prose/v2"
)


type Job struct {
	Content string
	Speaker string
	Id	string
}

func DivideTest(text string) []string {
	doc, _ := prose.NewDocument(text)
	
	sents := doc.Sentences()
	
	texts := []string{}

	for _, sent := range sents {
		texts = append(texts, sent.Text)
	}
	return texts
}