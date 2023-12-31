package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
)

func (hdl *HlsController) ServeJobList(w http.ResponseWriter, r *http.Request) {
	
	limit := 2
	offset := 0

	if r.URL.Query().Has("offset") {
		str := r.URL.Query().Get("offset")
		q, err := strconv.Atoi(str)
		if err != nil {
			offset = 0
		} else {
			offset = q
		}
	}

	if r.URL.Query().Has("limit") {
		str := r.URL.Query().Get("limit")
		q, err := strconv.Atoi(str)
		if err != nil {
			limit = 0
		} else {
			limit = q
		}
	}


	jobs, err := hdl.db.GetCompleteJob(limit, offset)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	 
	b, err := json.Marshal(jobs)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}