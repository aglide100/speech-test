package controller

import (
	"html/template"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"go.uber.org/zap"
)


func (hdl *HlsController) ServePlaylistFile(w http.ResponseWriter, r *http.Request) {
	const playlist =`#EXTM3U
#EXT-X-VERSION:3
#EXT-X-MEDIA-SEQUENCE:0
#EXT-X-TARGETDURATION: {{.MaxDuration}}

{{range .Audio}}
#EXT-X-DISCONTINUITY
#EXTINF: {{.Duration}},
/data/{{.Name}}
{{end}}

#EXT-X-ENDLIST`

	logger.Debug("uri", zap.Any("r", r.RequestURI))

	jobId := -1

	if r.URL.Query().Has("jobId") {
		str := r.URL.Query().Get("jobId")

		fileExtension := filepath.Ext(str)
	
		if fileExtension != ".m3u8" {
			http.Error(w, "wrong approach", http.StatusInternalServerError)
			return
		}

		fileName := filepath.Base(str)
		name := strings.Replace(fileName, ".m3u8", "", -1)

		q, err := strconv.Atoi(name)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
        	return
		} 
		
		jobId = q
	}

	res, err := hdl.db.GetAudioIds(jobId)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	
	if len(res) == 0 {
		http.Error(w, "invalid job", http.StatusBadRequest)
		return	
	}

	max := 0.0
	for _, val := range res {
		// logger.Info("Float", zap.Any("f", val.Duration))
		max = math.Max(max, float64(val.Duration))

	}

    // logger.Info("jobId", zap.Any("j", jobId))

	t := template.Must(template.New("m3u8").Parse(playlist))

	data := map[string]interface{} {
		"MaxDuration" : max,
		"Audio": res,
	}

	w.Header().Set("Content-Type", "application/x-mpegurl")
	
	t.Execute(w, data)
}