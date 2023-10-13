package job

import (
	"database/sql"
	"reflect"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/google/uuid"
	"github.com/jdkato/prose/v2"
	"go.uber.org/zap"
)

type RequestQueue struct {
	reqs []*Request
	db *sql.DB
}

type Request struct {
	Text string
	Jobs []*Job
	Audio [][]byte
}

type Job struct {
	Content string
	Speaker string
	Id	string
}

func NewRequestQueue(db *sql.DB) *RequestQueue {
	return &RequestQueue{db: db}
}

func (req *RequestQueue) AddRequest(newReq *Request) {
	req.reqs = append(req.reqs, newReq)
}

func (req *RequestQueue) AddAudioInRequest(job *Job, audio []byte) bool {
	for idx1, request := range req.reqs {
		for idx2, val := range request.Jobs {

			logger.Info("jobs", zap.Any("job", val))
			if reflect.DeepEqual(val, job) {
				req.reqs[idx1].Audio[idx2] = audio
				logger.Info("Add audio", zap.Any("len",len(audio)))
				break
			}
		}
	}

	return req.CheckComplete()
}

func (req *RequestQueue) SaveToAudio(request *Request) {
	// TODO
	logger.Info("TODO; save audio", zap.Any("req", request))
}

func (req *RequestQueue) CheckComplete() bool {
	found := false
	index := -1
	for idx1, request := range req.reqs {
		ok := true
		for idx2, _ := range request.Jobs{
			if len(req.reqs[idx1].Audio[idx2]) == 0 {
				ok = false
				break
			}
		}

		if ok {
			found = true
			index = idx1
			break
		}
	}

	if found {
		tmp := req.reqs[index]
		req.SaveToAudio(tmp)

		if len(req.reqs) == 1 {
			req.reqs = []*Request{}
			return true
		}

		if index == len(req.reqs)-1 {
			req.reqs = req.reqs[:index]
			return true
		} 

		req.reqs = append(req.reqs[:index], req.reqs[index+1])

		return true
	}
	logger.Info("size", zap.Any("len", len(req.reqs)))

	return false
}

func DivideTest(text string, speaker string) *Request {
	doc, _ := prose.NewDocument(text)
	
	sents := doc.Sentences()
	
	request := &Request{
		Text: text,
	}

	jobs := []*Job{}
	for _, sent := range sents {
		newJob := &Job{
			Content: sent.Text,
			Speaker: speaker,
			Id: uuid.New().String(),
		}

		jobs = append(jobs, newJob)
	}
	request.Jobs = jobs

	return request
}