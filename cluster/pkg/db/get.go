package db

import (
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"go.uber.org/zap"
)

func (db *Database) GetTextId(text, speaker string) (int, error) {
	const q = `
	SELECT id FROM text
		WHERE value = ? AND speaker = ?
	`

	var textId int
	err := db.conn.QueryRow(q, text, speaker).Scan(&textId)
	if err != nil {
		return -1, err
	}	

	return textId, nil
}

func (db *Database) GetIncompleteJob() ([]request.Request, error) {
	const q = `
	SELECT j.id AS job_id, j.speaker AS job_speaker, GROUP_CONCAT(t.value ORDER BY jt.no SEPARATOR ' ') AS text
	FROM job j
	         LEFT JOIN job_text jt ON j.id = jt.job_id
	         LEFT JOIN text t ON jt.text_id = t.id
	         LEFT JOIN (
	    SELECT text_id, COUNT(*) AS audio_count
	    FROM audio
	    GROUP BY text_id
	) AS a ON t.id = a.text_id
	GROUP BY j.id, j.speaker, j.max_index
	HAVING SUM(a.audio_count) != j.max_index OR SUM(a.audio_count) IS NULL
	`

	rows, err := db.conn.Query(q)
	if err != nil {
		return nil, err
	}

	var reqs []request.Request  

	for rows.Next() {
		var req request.Request

		if err := rows.Scan(&req.Id, &req.Speaker, &req.Text); err != nil {
			return nil, err
		}
		logger.Info("debug", zap.String("req.Text", req.Text))

		reqs = append(reqs, req)
	}

	return reqs, nil
}