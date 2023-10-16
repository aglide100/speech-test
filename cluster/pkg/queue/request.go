package queue

import (
	"reflect"

	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"go.uber.org/zap"
)

type RequestQueue struct {
	reqs []*request.Request
}


func NewRequestQueue() *RequestQueue {
	return &RequestQueue{}
}

func (req *RequestQueue) AddRequest(newReq *request.Request) {
	req.reqs = append(req.reqs, newReq)
}

func (req *RequestQueue) AddAudioInRequest(job *job.Job, audio []byte) (bool, *request.Request) {
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

func (req *RequestQueue) CheckComplete() (bool, *request.Request) {
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

		if len(req.reqs) == 1 {
			req.reqs = []*request.Request{}
			return true, tmp
		}

		if index == len(req.reqs)-1 {
			req.reqs = req.reqs[:index]
			return true, tmp
		} 

		req.reqs = append(req.reqs[:index], req.reqs[index+1])

		return true, tmp
	}
	logger.Info("size", zap.Any("len", len(req.reqs)))

	return false, nil
}
