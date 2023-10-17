package db

import (
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"go.uber.org/zap"
)

func (db *Database) SaveAudio(parent int, req *request.Request) error {
	const q = `
	INSERT INTO Audio (Audio.parent, Audio.data, Audio.order) VALUES ($1, $2, $3)
	`
	for idx, val := range req.Audio {
		_, err := db.conn.Exec(q, parent, val, idx)
		if err != nil {
			logger.Error("Can't insert Audio", zap.Error(err))
			return err
		}
	}

	return nil
}


func (db *Database) SaveJob(req *request.Request) error {
	const q = `
	INSERT INTO Job (Job.text, Job.date, Job.max_index) VALUES ($1, now(), $2)
	`

	_, err := db.conn.Exec(q, req.Text, len(req.Jobs))
	if err != nil {
		return err
	}

	return nil 
}