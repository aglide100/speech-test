package audio

import (
	"context"
	"errors"
	"log"
	"time"

	pb_svc_audio "github.com/aglide100/speech-test/cluster/pb/svc/audio"
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/queue"
	"github.com/aglide100/speech-test/cluster/pkg/runner"
)

type AudioSrv struct {
	pb_svc_audio.AudioServiceServer
	token string
	q *queue.PriorityJobQueue
}

func NewAudioServiceServer(q *queue.PriorityJobQueue, token string) *AudioSrv {
	return &AudioSrv{
		q:q,
		token: token,
	}
}

func (s *AudioSrv) GenerateAudio(ctx context.Context, in *pb_svc_audio.Requirement) (*pb_svc_audio.Audio, error) {

	return nil, nil
}

func (s *AudioSrv) MakingNewJob(ctx context.Context, in *pb_svc_audio.Request) (*pb_svc_audio.Result, error) {
	if (in.Auth.Token != s.token) {
		log.Printf("From : %s", in.Auth.From)
		return &pb_svc_audio.Result{}, errors.New("invalid token")
	}

	request := job.DivideTest(in.Content, in.Speaker)

	if s.q.Size() + len(request.Jobs) > s.q.Length() {
		log.Printf("Current queue size : %d / request size : %d", s.q.Size(), len(request.Jobs))
		return &pb_svc_audio.Result{Error: "queue is full"}, errors.New("queue is full")
	}

	for _, job := range request.Jobs {
		newAllocate := &queue.Allocate{
			Job: *job,
		}
		
		s.q.Push(newAllocate)
	}

	return &pb_svc_audio.Result{}, nil
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
	}, nil
}

func (s *AudioSrv) SendingResult(ctx context.Context, in *pb_svc_audio.Audio) (*pb_svc_audio.Job, error) {

	return nil, nil
}