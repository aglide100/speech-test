package controller

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"go.uber.org/zap"
)


func (hdl *HlsController) ServePlaylist(w http.ResponseWriter, r *http.Request) {
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

	parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 3 {
        http.Error(w, "invalid url", http.StatusBadRequest)
        return
    }

    jobId := strings.Replace(parts[2], ".m3u8", "", -1)

	idx, err := strconv.Atoi(jobId)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
        return
	}

	res, err := hdl.db.GetAudioIds(idx)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	
	if len(res) == 0 {
		http.Error(w, "invalid job", http.StatusBadRequest)
		return	
	}

	max := 0.0
	for _, val := range res {
		logger.Info("Float", zap.Any("f", val.Duration))
		max = math.Max(max, float64(val.Duration))
		max = math.Round(max)+1
	}

    logger.Info("jobId", zap.Any("j", jobId))

	t := template.Must(template.New("m3u8").Parse(playlist))

	data := map[string]interface{} {
		"MaxDuration" : max,
		"Audio": res,
	}

	w.Header().Set("Content-Type", "application/x-mpegurl")
	
	t.Execute(w, data)

}