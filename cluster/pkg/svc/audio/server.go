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

	return nil, nil
}

func (s *AudioSrv) MakingNewJob(ctx context.Context, in *pb_svc_audio.MakingNewJobReq) (*pb_svc_audio.Error, error) {
	if (in.Auth.Token != s.token) {
		log.Printf("From : %s", in.Auth.Who)
		return &pb_svc_audio.Error{Msg: "invalid token"}, errors.New("invalid token")
	}

	req := request.MakeRequest(in.Content, in.Speaker)

	err := s.db.SaveJob(&request.Request{
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

	return &pb_svc_audio.CheckingJobRes{
		Job: &pb_svc_audio.Job{
			Content:  p.Value.Job.Content,
			Speaker:  p.Value.Job.Speaker,
			Id:  p.Value.Job.Id,
		},
	}, nil
}

func (s *AudioSrv) SendingResult(ctx context.Context, in *pb_svc_audio.SendingResultReq) (*pb_svc_audio.Error, error) {
	if in.Auth.Token != s.token { 
		log.Printf("From : %s", in.Auth.Who)
		return &pb_svc_audio.Error{Msg: "invalid token"}, errors.New("invalid token")
	}
	s.mu.Lock()
	ok, result := s.requests.AddAudioInRequest(&job.Job{
		Content: in.Job.Content,
		Speaker: in.Job.Speaker,
		Id: in.Job.Id,
	}, in.Audio.Data)
	
	s.mu.Unlock()
	if ok {
		parent, err := s.db.GetParent(result.Text)
		if err != nil {
			return &pb_svc_audio.Error{
				Msg: "Internal error",
			}, err
		}

		err = s.db.SaveAudio(parent, result)
		if err != nil {
			return &pb_svc_audio.Error{
				Msg: "Internal error",
			}, err
		}
		
		return &pb_svc_audio.Error{}, nil 
	}
	
	return &pb_svc_audio.Error{Msg: "Not complete"}, nil
}

func (s *AudioSrv) AddRequestInQueue(req *request.Request) error {
	logger.Info("Added", zap.Any("req", req))
	req.Audio = make([][]byte, len(req.Jobs))
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
		req := request.MakeRequest(val.Text, val.Speaker)

		err := s.AddRequestInQueue(req)
		if err != nil {
			return err
		}
	
	}

	return nil 
}