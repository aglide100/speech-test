package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
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

	jobs, err := hdl.db.GetCompleteJob(limit, offset)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}


	b, err := json.Marshal(jobs)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}