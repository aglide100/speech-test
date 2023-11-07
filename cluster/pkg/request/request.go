package request

import (
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/txt"
)

type Request struct {
	FullText string
	Speaker string
	Title string
	Jobs []*job.Job
	JobId int
}

func MakeRequest(text string, speaker string) (*Request, error) {
	request := &Request{}

	jobs := []*job.Job{}

	
	texts, err := txt.DivideText(text)
	if err != nil {
		return nil, err
	}

	for idx, sent := range texts {
		newJob := &job.Job{
			Content: sent,
			Speaker: speaker,
			// Id: uuid.New().String(),
			No: idx,
		}

		jobs = append(jobs, newJob)
	}

	request.Jobs = jobs
	request.FullText = text
	request.Speaker = speaker
	
	return request, nil
}
