package db

import (
	"database/sql"
	"strconv"

	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"go.uber.org/zap"
)

func (db *Database) SaveAudio(textId int, audio []byte, sec float32, speaker string) error {
	const q = `
	INSERT INTO audio(data, speaker, text_id, sec, written_date)
   		VALUES (?, ?, ?, ?, now())
	`

	logger.Info("SaveAudio", zap.Any("speaker", speaker), zap.Any("textId", textId), zap.Any("sec", sec))
	_, err := db.conn.Exec(q, audio, speaker, textId, sec)
	if err != nil {
		logger.Error("Can't insert Audio", zap.Error(err))
		return err
	}

	return nil
}


func (db *Database) SaveJob(req *request.Request) (*request.Request, error) {
	const q = `
	INSERT INTO job(date, max_index, speaker, title)
    	VALUES (now(), ?, ?, ?)
	`
	tx, err := db.conn.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	res, err := db.conn.Exec(q, len(req.Jobs), req.Speaker, req.Title)
	if err != nil {
		logger.Error("Can't Insert Job", zap.Error(err))
		return nil, err
	}

	jobId, err := res.LastInsertId()
	if err != nil {
		logger.Error("last insertId", zap.Error(err))
		return nil, err
	}

	var newReq = req
	

	for idx, job := range newReq.Jobs {
		textId, err := db.SaveText(job.Content, job.Speaker, int(jobId), idx)
		if err != nil {
			logger.Error("Can't save text")
			return nil, err
		}

		newReq.Jobs[idx].TextId = strconv.Itoa(int(textId))
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	
	return newReq, nil 
}

func (db *Database) SaveText(text, speaker string, jobId, order int) (int64, error) {
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
		return -1, err
	}

	defer tx.Rollback()

	textId, err := db.GetTextId(text, speaker)
	if err != nil {
		if err == sql.ErrNoRows {
			res, err2 := db.conn.Exec(q1, text, speaker) 
			if err2 != nil {
				logger.Error("Can't exec q1")
				return -1, err2
			}

			textId, err2 := res.LastInsertId()
			if err2 != nil {
				logger.Error("Can't exec q2 in lastInsertId")
				return -1, err2
			}

			_, err2 = db.conn.Exec(q2, jobId, textId, order)
			if err2 != nil {
				logger.Error("Can't exec q2")
				return -1, err2
			}

			return textId, nil
		}

		logger.Error("Can't get textId", zap.Error(err))
		return -1, err
	}

	_, err = db.conn.Exec(q2, jobId, textId, order)
	if err != nil {
		logger.Error("Can't query q2")
		return -1, err 
	}
	
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return int64(textId), nil 
}