package request

import (
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/google/uuid"
)

type Request struct {
	Text string
	Speaker string
	Jobs []*job.Job
	Id int
}

func MakeRequest(text string, speaker string) *Request {
	request := &Request{}

	jobs := []*job.Job{}

	texts := job.DivideTest(text)

	for idx, sent := range texts {
		newJob := &job.Job{
			Content: sent,
			Speaker: speaker,
			Id: uuid.New().String(),
			No: idx,
		}

		jobs = append(jobs, newJob)
	}

	request.Jobs = jobs
	request.Text = text
	request.Speaker = speaker
	
	return request
}
