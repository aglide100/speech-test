package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aglide100/speech-test/cluster/pkg/controller"
	"github.com/aglide100/speech-test/cluster/pkg/db"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)


func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%v", err)
		os.Exit(1)
	}
}

func realMain() error {
	const port = 9090
	
	myDB, err := db.NewDB()
	if err != nil {
		return err
	}
	
	c := cache.New(5*time.Minute, 10*time.Minute)

	ctl := controller.NewHlsController(myDB, c)

	http.HandleFunc("/playlist/", addHeaders(http.HandlerFunc(ctl.ServePlaylist)))
	http.HandleFunc("/", addHeaders(http.HandlerFunc(ctl.FileHandler)))
	
	logger.Info("Starting server on ", zap.Any("port", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	return nil
}


func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return
		}

		h.ServeHTTP(w, r)
	}
}

