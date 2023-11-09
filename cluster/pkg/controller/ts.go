package controller

import (
	"net/http"
	"strconv"
	"strings"

	"path/filepath"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

func (hdl *HlsController) ServeTsFile(w http.ResponseWriter, r *http.Request) {
	fileExtension := filepath.Ext(r.RequestURI)
	
	logger.Debug("r", zap.Any("r", r.RequestURI))

	if fileExtension != ".ts" {
		http.Error(w, "wrong approach", http.StatusInternalServerError)
		return
	}

	fileName := filepath.Base(r.RequestURI)
	logger.Info("filename", zap.Any("name", fileName))

	name := strings.Replace(fileName, ".ts", "", -1)

	res := strings.Split(name, "_")

	textId := res[0]
	// no := res[1]


	idx, err := strconv.Atoi(textId)
	if err != nil {
		http.Error(w, "check file name", http.StatusInternalServerError)
		return
	}

	value, ok := hdl.c.Get(name)
	if ok {
		ba, ok := value.([]byte)
		if ok {
			_, err = w.Write(ba)
			if err != nil {
				http.Error(w, "Failed to write response", http.StatusInternalServerError)
				return
			}
		}
	}

	data, err := hdl.db.GetAudio(idx)
	hdl.c.Set(name, data, cache.DefaultExpiration)
	if err != nil {
		logger.Error("Can't get audio", zap.Error(err), zap.Any("idx", idx))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}