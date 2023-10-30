package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aglide100/speech-test/cluster/pkg/db"
)


var myDB *db.Database
func main() {
	if err := realMain(); err != nil {
		log.Printf("err :%v", err)
		os.Exit(1)
	}
}

func realMain() error {
	const mediaDir = "output"
	const port = 9090
	
	conn, err := db.NewDB()
	if err != nil {
		return err
	}

	myDB = conn



	http.HandleFunc("/byte", addHeaders(http.HandlerFunc(byteHandler)))
	http.Handle("/", addHeaders(http.FileServer(http.Dir(mediaDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving on : %d\n", port)

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

func byteHandler(w http.ResponseWriter, r *http.Request) {
	data, err := myDB.GetAudio(34)
	if err != nil {
		log.Panicf("err ", err.Error())
	}

	// w.Header().Set("Content-Type", "application/octet-stream")
	
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}