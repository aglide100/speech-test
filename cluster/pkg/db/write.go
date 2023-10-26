package db

import (
	"database/sql"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"go.uber.org/zap"
)

func (db *Database) SaveAudio(textId int, audio []byte, millisec float32, speaker string) error {
	const q = `
	INSERT INTO audio(data, speaker, text_id, millisec)
   		VALUES (?, ?, ?, ?)
	`

	logger.Info("SaveAudio")
	_, err := db.conn.Exec(q, audio, speaker, textId, millisec)
	if err != nil {
		logger.Error("Can't insert Audio", zap.Error(err))
		return err
	}

	return nil
}


func (db *Database) SaveJob(req *request.Request) error {
	const q = `
	INSERT INTO job(date, max_index, speaker)
    	VALUES (now(), ?, ?)
	`
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	res, err := db.conn.Exec(q, len(req.Jobs), req.Speaker)
	if err != nil {
		logger.Error("Can't Insert Job", zap.Error(err))
		return err
	}

	jobId, err := res.LastInsertId()
	if err != nil {
		logger.Error("last insertId", zap.Error(err))
		return err
	}

	for idx, job := range req.Jobs {
		err = db.SaveText(job.Content, job.Speaker, int(jobId), idx)
		if err != nil {
			logger.Error("Can't save text")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	
	return nil 
}

func (db *Database) SaveText(text, speaker string, jobId, order int) error {
	const q1 = `
	INSERT INTO text(value, speaker)
		VALUES (?, ?)
	`

	const q2 = `
	INSERT INTO job_text(job_id, text_id, no)
    	VALUES (?, ?, ?)
	`

	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	textId, err := db.GetTextId(text, speaker)
	if err != nil {
		if err == sql.ErrNoRows {
			res, err2 := db.conn.Exec(q1, text, speaker) 
			if err2 != nil {
				logger.Error("Can't exec q1")
				return err2
			}

			textId, err2 := res.LastInsertId()
			if err2 != nil {
				logger.Error("Can't exec q2 in lastInsertId")
				return err2
			}

			_, err2 = db.conn.Exec(q2, jobId, textId, order)
			if err2 != nil {
				logger.Error("Can't exec q2")
				return err2
			}

			return nil
		}

		logger.Error("Can't get textId", zap.Error(err))
		return err
	}

	_, err = db.conn.Exec(q2, jobId, textId, order)
	if err != nil {
		logger.Error("Can't query q2")
		return err 
	}
	
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil 
}