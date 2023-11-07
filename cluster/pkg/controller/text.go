package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"go.uber.org/zap"
)


type Text struct {
	Txt string
}

func (hdl *HlsController) ServeText(w http.ResponseWriter, r *http.Request) {

	logger.Debug("r", zap.Any("r", r.RequestURI))

	jobId := -1

	if r.URL.Query().Has("jobId") {
		str := r.URL.Query().Get("jobId")
		q, err := strconv.Atoi(str)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
        	return
		} 
		
		jobId = q
	}

	res, err := hdl.db.GetTextFromJob(jobId)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	
	b, err := json.Marshal(Text{
		Txt: res,
	})
	
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