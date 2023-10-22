package db

import (
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"go.uber.org/zap"
)

func (db *Database) SaveAudio(parent int, req *request.Request) error {
	const q = `
	INSERT INTO Audio (Audio.parent, Audio.data, Audio.no) VALUES (?, ?, ?)
	`
	for idx, val := range req.Audio {
		_, err := db.conn.Exec(q, parent, string(val[:]), idx)
		if err != nil {
			logger.Error("Can't insert Audio", zap.Error(err))
			return err
		}
	}

	return nil
}


func (db *Database) SaveJob(req *request.Request) error {
	const q = `
	INSERT INTO speech.Job (date, max_index, speaker) VALUES (now(), ?, ?)	
	`

	_, err := db.conn.Exec(q, len(req.Jobs), req.Speaker)
	if err != nil {
		return err
	}

	const q2 = `
	SELECT LAST_INSERT_ID()
	`

	var parentId int

	err = db.conn.QueryRow(q2).Scan(&parentId)
	if err != nil {
		return err
	}

	for idx, job := range req.Jobs {
		err = db.SaveText(job.Content, parentId, idx)
		if err != nil {
			return err
		}
	}
	
	return nil 
}

func (db *Database) SaveText(text string, id, order int) error {
	const q = `
	INSERT INTO speech.Texts (Texts.value, Texts.parent, Texts.no) VALUES (?, ?, ?)
	`

	_, err := db.conn.Exec(q, text, id, order)
	if err != nil {
		return err
	}

	return nil 
}