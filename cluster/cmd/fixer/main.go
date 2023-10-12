package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	pb_svc_audio "github.com/aglide100/speech-test/cluster/pb/svc/audio"

	"github.com/aglide100/speech-test/cluster/pkg/queue"
	"github.com/aglide100/speech-test/cluster/pkg/svc/audio"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)
var (
	managerAddr = flag.String("manager grpc addr", "0.0.0.0:50012", "grpc address")
	token = flag.String("token", "test", "using for auth")
	password = flag.String("password", "password", "using for db password")
)

func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%v", err)
		os.Exit(1)
	}
}

func realMain() error {
	gprcListener, err := net.Listen("tcp", fmt.Sprintf(*managerAddr))
	if err != nil {
		return err
	}
	defer gprcListener.Close()

	var wait sync.WaitGroup
	wait.Add(1)
	queue := queue.NewJobQueue()

	audioSrv := audio.NewAudioServiceServer(queue, *token)
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb_svc_audio.RegisterAudioServiceServer(grpcServer, audioSrv)

	wg, _ := errgroup.WithContext(context.Background())

	wg.Go(func() error {
		log.Printf("Starting grpcServer at: %s" ,*managerAddr)
		err := grpcServer.Serve(gprcListener)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}

		return nil
	})

	wg.Go(func() error {
		duration, _ := time.ParseDuration("1s")
		ticker := time.NewTicker(duration)
		defer ticker.Stop()


		for range ticker.C {
			queue.CheckTimeOut()
		}

		return nil
	})

	return wg.Wait()
}