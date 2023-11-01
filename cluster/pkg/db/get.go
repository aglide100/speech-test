package db

import (
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (db *Database) GetAudio(textId int) ([]byte, error) {
	const q = `
	SELECT data
	FROM audio AS a
	WHERE a.text_id = ?
	`

	var data []byte

	err := db.conn.QueryRow(q, textId).Scan(&data)
	if err != nil {
		return nil, err
	}
	
	logger.Debug("info", zap.Any("data", data))
	return data, err
}

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


func (db *Database) GetJobText(jobId int) (job.Job, error) {
	const q = `
	SELECT j.id, GROUP_CONCAT(t.value ORDER BY jt.no) AS text
	FROM job AS j
	    LEFT JOIN (
	        SELECT job_id, text_id, no
	        FROM job_text
	        WHERE no = 0 OR no =1
	    ) AS jt ON jt.job_id = j.id
	    LEFT JOIN text AS t ON t.id = jt.text_id
	WHERE j.id = ?
	`

	var j job.Job
	err := db.conn.QueryRow(q, jobId).Scan(&j.Id, &j.Content)
	if err != nil {
		return j, err
	}
	
	return j, nil
}

func (db *Database) GetAudioIds(jobId int) ([]job.Audio, error) {
	const q = `
	SELECT a.text_id AS Name, a.sec AS Duration, jt.no
	FROM audio AS a 
		LEFT JOIN job_text jt ON a.text_id = jt.text_id
	WHERE jt.job_id = ?
	ORDER BY jt.no
	`

	rows, err := db.conn.Query(q, jobId)
	if err != nil {
		return nil, err
	}
	var data []job.Audio
	
	for rows.Next() {
		var a job.Audio

		if err := rows.Scan(&a.Name, &a.Duration, &a.No); err != nil {
			return nil, err
		}
		a.Name = a.Name + ".ts"
		a.Duration = float32(float64(a.Duration / 1000))

		data = append(data, a)
	}

	return data, nil
}


func (db *Database) GetIncompleteJob() ([]request.Request, error) {
	const q = `
	SELECT j.id AS job_id, j.speaker AS job_speaker, GROUP_CONCAT(t.value ORDER BY jt.no SEPARATOR ' ') AS text
	FROM job j
	    LEFT JOIN job_text AS jt ON j.id = jt.job_id
	    LEFT JOIN text AS t ON jt.text_id = t.id
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

		reqs = append(reqs, req)
	}

	return reqs, nil
}

func (db *Database) GetCompleteJob(limit, offset int) ([]*job.Job, error) {
	const q = `
	SELECT j.id      AS job_id,
       j.speaker AS job_speaker
	FROM job j
	    LEFT JOIN job_text jt ON j.id = jt.job_id
	    LEFT JOIN (SELECT text_id, COUNT(*) AS audio_count
	        FROM audio
	        GROUP BY text_id
		) AS a ON jt.text_id = a.text_id
	GROUP BY j.id, j.speaker, j.max_index
	HAVING SUM(a.audio_count) = j.max_index
	    OR SUM(a.audio_count) IS NOT NULL
	LIMIT ? OFFSET ?
	`

	rows, err := db.conn.Query(q, limit, offset)
	if err != nil {
		return nil, err
	}

	var jobs []*job.Job  

	for rows.Next() {
		var tmp job.Job

		if err := rows.Scan(&tmp.Id, &tmp.Speaker); err != nil {
			return nil, err
		}

		newJob := &job.Job{
			Content: tmp.Content,
			Speaker: tmp.Speaker,
			No: tmp.No,
			Id: uuid.New().String(),
		}
		
		jobs = append(jobs, newJob)
	}

	return jobs, nil
}

func (db *Database) GetIncompleteAudio(jobId int) ([]*job.Job, error) {
	const q = `
	SELECT t.value, t.speaker, jt.no
	FROM text AS t
	    LEFT JOIN audio AS a ON t.id = a.text_id
	    LEFT JOIN job_text AS jt ON t.id = jt.text_id
	WHERE a.text_id IS NULL AND jt.job_id = ?
	`

	rows, err := db.conn.Query(q, jobId)
	if err != nil {
		return nil, err
	}

	var jobs []*job.Job  

	for rows.Next() {
		var tmp job.Job

		if err := rows.Scan(&tmp.Content, &tmp.Speaker, &tmp.No); err != nil {
			return nil, err
		}

		newJob := &job.Job{
			Content: tmp.Content,
			Speaker: tmp.Speaker,
			No: tmp.No,
			Id: uuid.New().String(),
		}
		
		jobs = append(jobs, newJob)
	}

	return jobs, nil
}