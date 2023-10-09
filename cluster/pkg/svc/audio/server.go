package audio

import (
	"context"
	"errors"
	"log"
	"time"

	pb_svc_audio "github.com/aglide100/speech-test/cluster/pb/svc/audio"
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/queue"
	"github.com/aglide100/speech-test/cluster/pkg/runner"
	"go.uber.org/zap"
)

type AudioSrv struct {
	pb_svc_audio.AudioServiceServer
	token string
	q *queue.PriorityJobQueue
	requests *job.RequestQueue
}

func NewAudioServiceServer(q *queue.PriorityJobQueue, token string) *AudioSrv {
	return &AudioSrv{
		q:q,
		token: token,
		requests: job.NewRequestQueue(),
	}
}

func (s *AudioSrv) GenerateAudio(ctx context.Context, in *pb_svc_audio.Requirement) (*pb_svc_audio.Audio, error) {

	return nil, nil
}

func (s *AudioSrv) MakingNewJob(ctx context.Context, in *pb_svc_audio.Request) (*pb_svc_audio.Error, error) {
	if (in.Auth.Token != s.token) {
		log.Printf("From : %s", in.Auth.From)
		return &pb_svc_audio.Error{Msg: "invalid token"}, errors.New("invalid token")
	}

	request := job.DivideTest(in.Content, in.Speaker)
	request.Audio = make([][]byte, len(request.Jobs))
	s.requests.AddRequest(request)

	for _, job := range request.Jobs {
		newAllocate := &queue.Allocate{
			Job: *job,
		}
		
		s.q.Push(newAllocate)
	}

	return &pb_svc_audio.Error{}, nil
}

func (s *AudioSrv) CheckingJob(ctx context.Context, in *pb_svc_audio.Checking) (*pb_svc_audio.Job, error) {
	if (in.Auth.Token != s.token) {
		log.Printf("From : %s", in.Auth.From)
		return &pb_svc_audio.Job{}, errors.New("invalid token")
	}

	job, found := s.q.GetNotAllocate()
	if !found {
		return &pb_svc_audio.Job{}, errors.New("there's no available jobs")
	}

	allocated := &queue.Allocate{
		Job: job,
		Who: runner.Runner{
			CurrentWork: job.Content,
			Who: in.Auth.From,
		},
		When: time.Now(),
	}

	s.q.SetAllocate(allocated)

	return &pb_svc_audio.Job{
		Content: job.Content,
		Speaker: job.Speaker,
		Id: job.Id,
	}, nil
}

func (s *AudioSrv) SendingResult(ctx context.Context, in *pb_svc_audio.AudioResult) (*pb_svc_audio.Error, error) {
	if in.Auth.Token != s.token { 
		log.Printf("From : %s", in.Auth.From)
		return &pb_svc_audio.Error{Msg: "invalid token"}, errors.New("invalid token")
	}
	logger.Info("audio", zap.Any("bytes", in.Audio.Data))

	ok := s.requests.AddAudioInRequest(&job.Job{
		Content: in.Job.Content,
		Speaker: in.Job.Speaker,
		Id: in.Job.Id,
	}, in.Audio.Data)

	if ok {
		return &pb_svc_audio.Error{}, nil 
	}

	return &pb_svc_audio.Error{Msg: "Not complete"}, nil
}