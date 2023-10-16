package request

import (
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/google/uuid"
)

type Request struct {
	Text string
	Jobs []*job.Job
	Audio [][]byte
}

func MakeRequest(text string, speaker string) *Request {
	request := &Request{}

	jobs := []*job.Job{}

	texts := job.DivideTest(text)
	
	for _, sent := range texts {
		newJob := &job.Job{
			Content: sent,
			Speaker: speaker,
			Id: uuid.New().String(),
		}

		text += sent

		jobs = append(jobs, newJob)
	}

	
	request.Jobs = jobs
	request.Text = text
	
	return request
}