package audio

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	pb_svc_audio "github.com/aglide100/speech-test/cluster/pb/svc/audio"
	"github.com/aglide100/speech-test/cluster/pkg/db"
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/queue"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"github.com/aglide100/speech-test/cluster/pkg/runner"
	"go.uber.org/zap"
)

type AudioSrv struct {
	pb_svc_audio.AudioServiceServer
	token string
	running *queue.PriorityQueue
	waiting *queue.PriorityQueue
	requests *queue.RequestQueue
	mu *sync.Mutex
	db *db.Database
}

func NewAudioServiceServer(running, waiting *queue.PriorityQueue, token string, mutex *sync.Mutex, db *db.Database) *AudioSrv {
	return &AudioSrv{
		running: running,
		waiting: waiting, 
		token: token,
		requests: queue.NewRequestQueue(),
		mu: mutex,
		db: db,
	}
}

func (s *AudioSrv) GetAudio(ctx context.Context, in *pb_svc_audio.GetAudioReq) (*pb_svc_audio.Audio, error) {
	res, err := s.db.GetAudio(int(in.JobId))
	if err != nil {
		logger.Error("!", zap.Error(err))
	}




	return &pb_svc_audio.Audio{
		Data: res,
	}, nil
}

func (s *AudioSrv) MakingNewJob(ctx context.Context, in *pb_svc_audio.MakingNewJobReq) (*pb_svc_audio.Error, error) {
	if (in.Auth.Token != s.token) {
		log.Printf("From : %s", in.Auth.Who)
		return &pb_svc_audio.Error{Msg: "invalid token"}, errors.New("invalid token")
	}

	logger.Info("MakingNewJob", zap.Any("content",in.Content))
	req, err := request.MakeRequest(in.Content, in.Speaker)
	if err != nil {
		return &pb_svc_audio.Error{
			Msg: err.Error(),
		}, nil
	}

	err = s.db.SaveJob(&request.Request{
		Text: in.Content,
		Jobs: req.Jobs,
		Speaker: in.Speaker,
	})
	if err != nil {
		return nil, err
	}

	err = s.AddRequestInQueue(req)
	if err != nil {
		return nil, err
	}

	return &pb_svc_audio.Error{}, nil
}

func (s *AudioSrv) CheckingJob(ctx context.Context, in *pb_svc_audio.CheckingJobReq) (*pb_svc_audio.CheckingJobRes, error) {
	if (in.Auth.Token != s.token) {
		log.Printf("From : %s", in.Auth.Who)
		return &pb_svc_audio.CheckingJobRes{}, errors.New("invalid token")
	}

	s.mu.Lock()
	p, ok := s.waiting.Pop()

	if ok {
		allocated := &queue.Allocate{
			Job: job.Job{
				Content: p.Value.Job.Content,
				Speaker:  p.Value.Job.Speaker,
				Id:  p.Value.Job.Id,
				No: p.Value.Job.No,
			},
			Who: runner.Runner{
				CurrentWork:  p.Value.Job.Content,
				Who: in.Auth.Who,
			},
			When: time.Now(),
		}
	
		s.running.Push(&queue.Item{
			Value: *allocated,
		})
	}

	s.mu.Unlock()

	if !ok {
		return &pb_svc_audio.CheckingJobRes{
		}, nil
	}

	logger.Info("Pop in wating", zap.Any("who", in.Auth.Who), zap.Any("content", p.Value.Job.Content))

	return &pb_svc_audio.CheckingJobRes{
		Job: &pb_svc_audio.Job{
			Content:  p.Value.Job.Content,
			Speaker:  p.Value.Job.Speaker,
			Id:  p.Value.Job.Id,
			No: int32(p.Value.Job.No),
		},
	}, nil
}

func (s *AudioSrv) SendingResult(ctx context.Context, in *pb_svc_audio.SendingResultReq) (*pb_svc_audio.Error, error) {
	if in.Auth.Token != s.token { 
		log.Printf("From : %s", in.Auth.Who)
		return &pb_svc_audio.Error{Msg: "invalid token"}, errors.New("invalid token")
	}
	s.mu.Lock()

	ok, result := s.requests.RemoveJobInRequest(&job.Job{
		Content: in.Job.Content,
		Speaker: in.Job.Speaker,
		Id: in.Job.Id,
		No: int(in.Job.No),
	})
	
	if ok {
		textId, err := s.db.GetTextId(in.Job.Content, in.Job.Speaker)
		if err != nil {
			
			return &pb_svc_audio.Error{
				Msg: "Internal error",
			}, err
		}
	
		err = s.db.SaveAudio(textId, in.Audio.Data, in.Audio.Millisec, in.Job.Speaker)
		if err != nil {
			return &pb_svc_audio.Error{
				Msg: "Internal error",
			}, err
		}

		if len(result.Jobs) == 0 {
			ok := s.requests.RemoveRequest(result)
			if !ok {
				logger.Error("Can't remove request in queue!", zap.Any("req", s.requests))
			}

			err := s.db.UpdateTotalPlayingTime(result.Id)
			if err != nil {
				return &pb_svc_audio.Error{
					Msg: "Internal error",
				}, err
			}
		}

		s.mu.Unlock()
		return &pb_svc_audio.Error{Msg: "Done"}, nil 
	}

	s.mu.Unlock()
	return &pb_svc_audio.Error{Msg: "Not complete"}, nil
}

func (s *AudioSrv) AddRequestInQueue(req *request.Request) error {
	logger.Info("Added", zap.Any("req", req))
	s.requests.AddRequest(req)

	for _, job := range req.Jobs {
		newAllocate := queue.Allocate{
			Job: *job,
		}

		s.mu.Lock()

		item := &queue.Item{
			Value: newAllocate, 
			Index : s.waiting.Len(),
		}
		s.waiting.Push(item)

		s.mu.Unlock()
	}

	return nil
}

func (s *AudioSrv) AddIncomplete() error {
	res, err := s.db.GetIncompleteJob()
	if err != nil {
		return err
	}

	for _, val := range res {
		logger.Info("id", zap.Any("val", val))
		jobs, err := s.db.GetIncompleteAudio(val.Id)
		if err != nil {
			logger.Info("Can't add job in req", zap.Error(err))
			return err
		}


		req := &request.Request{
			Text: val.Text,
			Speaker: val.Speaker,
			Jobs: jobs,
		}
		// req, err := request.MakeRequest(val.Text, val.Speaker)
		// if err != nil {
		// 	return err
		// }

		err = s.AddRequestInQueue(req)
		if err != nil {
			return err
		}
	
	}

	return nil 
}