package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
)

func main() {

	logger.Debug("!!!")
	logger.Info("@@@")

	// if err := realMain(); err != nil {
	// 	log.Printf("err :%v", err)
	// 	os.Exit(1)
	// }
}

func realMain() error {
	const mediaDir = "output"
	const port = 8080

	http.Handle("/", addHeaders(http.FileServer(http.Dir(mediaDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving on : %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	return nil
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}