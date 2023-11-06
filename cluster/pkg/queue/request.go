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

func (req *RequestQueue) RemoveRequest(target *request.Request) bool {
	found := false;
	index := -1
	
	for idx, val := range req.reqs {
		if reflect.DeepEqual(val, target) {
			found = true
			index = idx
			break
		}
	}

	if found {
		if len(req.reqs) == 1 {
			req.reqs = make([]*request.Request, 0)
			return true
		}

		if index == len(req.reqs) -1 {
			req.reqs = append(req.reqs[:index])
			return true
		}

		req.reqs = append(req.reqs[:index], req.reqs[index+1])
		return true
	}

	return false
}

func (req *RequestQueue) RemoveJobInRequest(j *job.Job) (bool, *request.Request, bool) {
	logger.Debug("RemoveJobInRequest", zap.Any("job", j))
	found := false
	req_idx := -1
	job_idx := -1

	for idx1, req := range req.reqs {
		for idx2, val := range req.Jobs {
			if reflect.DeepEqual(val, j) {
				found = true
				req_idx = idx1
				job_idx = idx2
				break;
			}
		}	
	}

	if found {
		if len(req.reqs[req_idx].Jobs) == 1 {
			req.reqs[req_idx].Jobs = make([]*job.Job, 0)
			return true, req.reqs[req_idx], true
		}

		if job_idx == len(req.reqs[req_idx].Jobs)-1 {
			req.reqs[req_idx].Jobs = req.reqs[req_idx].Jobs[:job_idx]
			return true, req.reqs[req_idx], false
		}

		req.reqs[req_idx].Jobs = append(req.reqs[req_idx].Jobs[:job_idx], req.reqs[req_idx].Jobs[job_idx+1])
		return true, req.reqs[req_idx], false
	}
	
	return false, nil, false
}
