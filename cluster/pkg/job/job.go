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
	Speaker string
	Id	int
	IsComplete bool
}

func DivideTest(text string, speaker string) *Request {
	doc, _ := prose.NewDocument(text)
	
	sents := doc.Sentences()
	
	request := &Request{
		Text: text,
	}

	jobs := []*Job{}
	idx := 0
	for _, sent := range sents {
		newJob := &Job{
			Content: sent.Text,
			Speaker: speaker,
			Id: idx,
			IsComplete: false,
		}

		idx++

		jobs = append(jobs, newJob)
	}
	request.Jobs = jobs

	return request
}