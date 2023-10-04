package job

import (
	"github.com/jdkato/prose/v2"
)

type Request struct {
	Text string
	Jobs []*Job
}

type Job struct {
	Content string
	Id	int
}

func DivideTest(text string) *Request {
	doc, _ := prose.NewDocument(text)

	sents := doc.Sentences()
	

	request := &Request{
		Text: text,
	}

	jobs := []Job{}
	idx := 0
	for _, sent := range sents {
		newJob := &Job{
			Content: sent.Text,
			Id: idx,
		}

		idx++

		jobs = append(jobs, *newJob)
	}

	return request
}