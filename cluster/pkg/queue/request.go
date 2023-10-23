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

func (req *RequestQueue) RemoveJobInRequest(j *job.Job) (bool, *request.Request) {
	logger.Info("RemoveJobInRequest", zap.Any("job", j))
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
			return true, req.reqs[req_idx]
		}

		if job_idx == len(req.reqs[req_idx].Jobs)-1 {
			req.reqs[req_idx].Jobs = req.reqs[req_idx].Jobs[:job_idx]
			return true, req.reqs[req_idx]
		}

		req.reqs[req_idx].Jobs = append(req.reqs[req_idx].Jobs[:job_idx], req.reqs[req_idx].Jobs[job_idx+1])
		return true, req.reqs[req_idx]
	}

	logger.Info("Can't remove Job in Request", zap.Any("req", req.reqs))

	return false, nil
}

// func (req *RequestQueue) AddAudioInRequest(job *job.Job, audio []byte) (bool, *request.Request) {
	
	
	
// 	// for idx1, request := range req.reqs {
// 	// 	for idx2, val := range request.Jobs {

// 	// 		logger.Info("jobs", zap.Any("job", val))
// 	// 		if reflect.DeepEqual(val, job) {
// 	// 			req.reqs[idx1].Audio[idx2] = audio
// 	// 			logger.Info("Add audio", zap.Any("len",len(audio)))
// 	// 			break
// 	// 		}
// 	// 	}
// 	// }

// 	return req.CheckComplete()
// }

// func (req *RequestQueue) CheckComplete() (bool, *request.Request) {
// 	for idx1, request := range req.reqs {


// 	}

// 	// found := false
// 	// index := -1
// 	// for idx1, request := range req.reqs {
// 	// 	ok := true
// 	// 	for idx2, _ := range request.Jobs{
// 	// 		if len(req.reqs[idx1].Audio[idx2]) == 0 {
// 	// 			ok = false
// 	// 			break
// 	// 		}
// 	// 	}

// 	// 	if ok {
// 	// 		found = true
// 	// 		index = idx1
// 	// 		break
// 	// 	}
// 	// }

// 	// if found {
// 	// 	tmp := req.reqs[index]

// 	// 	if len(req.reqs) == 1 {
// 	// 		req.reqs = []*request.Request{}
// 	// 		return true, tmp
// 	// 	}

// 	// 	if index == len(req.reqs)-1 {
// 	// 		req.reqs = req.reqs[:index]
// 	// 		return true, tmp
// 	// 	} 

// 	// 	req.reqs = append(req.reqs[:index], req.reqs[index+1])

// 	// 	return true, tmp
// 	// }
// 	// logger.Info("size", zap.Any("len", len(req.reqs)))

// 	return false, nil
// }
