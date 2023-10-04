package audio

import (
	"context"
	"errors"
	"log"

	pb_svc_audio "github.com/aglide100/speech-test/cluster/pb/svc/audio"
	"github.com/aglide100/speech-test/cluster/pkg/queue"
)

type AudioSrv struct {
	pb_svc_audio.AudioServiceServer
	token string
	q *queue.JobQueue
}

func NewAudioServiceServer(q *queue.JobQueue, token string) *AudioSrv {
	return &AudioSrv{
		q:q,
		token: token,
	}
}

func (s *AudioSrv) GenerateAudio(ctx context.Context, in *pb_svc_audio.Requirement) (*pb_svc_audio.Audio, error) {

	return nil, nil
}

func (s *AudioSrv) MakingNewJob(ctx context.Context, in *pb_svc_audio.Request) (*pb_svc_audio.Audio, error) {
	if (in.Auth.Token != s.token) {
		log.Printf("From : %s", in.Auth.From)
		return nil, errors.New("invalid token")
	}

	
	return nil, nil
}

func (s *AudioSrv) CheckingJob(ctx context.Context, in *pb_svc_audio.Checking) (*pb_svc_audio.Job, error) {

	return nil, nil
}

func (s *AudioSrv) SendingResult(ctx context.Context, in *pb_svc_audio.Audio) (*pb_svc_audio.Job, error) {

	return nil, nil
}