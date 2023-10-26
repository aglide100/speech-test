package db

import (
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"go.uber.org/zap"
)

func (db *Database) UpdateTotalPlayingTime(jobId int) error {
	const q = `
	UPDATE job AS j
    	LEFT JOIN (
    	    SELECT jt.job_id, SUM(a.millisec) AS total_millisec
    	    FROM job_text AS jt
    	        LEFT JOIN audio AS a ON a.text_id = jt.text_id
    	    GROUP BY jt.job_id
    	) AS sub ON j.id = sub.job_id
	SET j.playing_time = IFNULL(sub.total_millisec, 0)
	WHERE j.id = ?;
	`

	_, err := db.conn.Exec(q, jobId)
	if err != nil {
		logger.Error("Can't update total playing time", zap.Error(err))
		return err
	}

	return nil
}