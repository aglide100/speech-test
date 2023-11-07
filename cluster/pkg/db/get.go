package db

import (
	"github.com/aglide100/speech-test/cluster/pkg/job"
	"github.com/aglide100/speech-test/cluster/pkg/logger"
	"github.com/aglide100/speech-test/cluster/pkg/request"
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


// func (db *Database) GetTextSummaryFromJob(jobId int) (job.Job, error) {
// 	const q = `
// 	SELECT j.id, GROUP_CONCAT(t.value ORDER BY jt.no) AS text
// 	FROM job AS j
// 	    LEFT JOIN (
// 	        SELECT job_id, text_id, no
// 	        FROM job_text
// 	        WHERE no = 0 OR no =1
// 	    ) AS jt ON jt.job_id = j.id
// 	    LEFT JOIN text AS t ON t.id = jt.text_id
// 	WHERE j.id = ?
// 	`

// 	var j job.Job
// 	err := db.conn.QueryRow(q, jobId).Scan(&j.Id, &j.Content)
// 	if err != nil {
// 		return j, err
// 	}
	
// 	return j, nil
// }

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

func (db *Database) GetIncompleteJobText(jobId int, speaker string) ([]job.Job, error) {
	const q = `
	SELECT jt.text_id AS text_id, t.value AS text, no
	FROM job_text AS jt
	    LEFT JOIN text AS t on jt.text_id = t.id
	    LEFT JOIN audio AS a ON a.text_id = jt.text_id
	WHERE job_id = ? AND a.data IS NULL
	`

	rows, err := db.conn.Query(q, jobId)
	if err != nil {
		return nil, err
	}

	var jts []job.Job
	for rows.Next() {
		var jt job.Job

		if err := rows.Scan(&jt.TextId, &jt.Content, &jt.No); err != nil {
			return nil, err
		}

		jt.Speaker = speaker

		jts = append(jts, jt)
	}

	return jts, nil
}

func (db *Database) GetIncompleteJobIDs() ([]request.Request, error) {
	const q = `
	SELECT j.id AS job_id, j.speaker AS job_speaker
	FROM job AS j
	         LEFT JOIN job_text AS jt ON j.id = jt.job_id
	         LEFT JOIN (
	    SELECT text_id, COUNT(*) AS audio_count
	    FROM audio
	    GROUP BY text_id
	) AS a ON jt.text_id = a.text_id
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

		if err := rows.Scan(&req.JobId, &req.Speaker); err != nil {
			return nil, err
		}

		reqs = append(reqs, req)
	}

	return reqs, nil
}

func (db *Database) GetTextFromJob(jobId int) (string, error) {
	const q = `
	SELECT GROUP_CONCAT(value)
	FROM job AS j
	    LEFT JOIN job_text AS jt ON j.id = jt.job_id
	    LEFT JOIN (
	        SELECT id, value
	        FROM text
	    ) AS t ON t.id = jt.text_id
	WHERE job_id = ?
	`

	var text string

	err := db.conn.QueryRow(q, jobId).Scan(&text)
	if err != nil {
		return "", err
	}

	return text, nil
}

func (db *Database) GetCompleteJob(limit, offset int) ([]*job.ReturningJob, error) {
	const q = `
	SELECT j.id  AS job_id,
       j.speaker AS job_speaker,
       j.title AS job_title,
       j.playing_time
	FROM job AS j
	         LEFT JOIN job_text AS jt ON j.id = jt.job_id
	         LEFT JOIN (SELECT text_id, COUNT(*) AS audio_count
	                    FROM audio
	                    GROUP BY text_id
	) AS a ON jt.text_id = a.text_id
	GROUP BY j.id, j.speaker, j.max_index
	HAVING SUM(a.audio_count) = j.max_index
	LIMIT ? OFFSET ?
	`

	rows, err := db.conn.Query(q, limit, offset)
	if err != nil {
		return nil, err
	}

	var jobs []*job.ReturningJob   

	for rows.Next() {
		var tmp job.ReturningJob

		if err := rows.Scan(&tmp.Id, &tmp.Speaker, &tmp.Title, &tmp.PlayingTime); err != nil {
			return nil, err
		}

		newJob := &job.ReturningJob{
			Id: tmp.Id,
			Content: tmp.Content,
			Speaker: tmp.Speaker,
			Title: tmp.Title,
			PlayingTime: tmp.PlayingTime,
		}
		
		jobs = append(jobs, newJob)
	}

	return jobs, nil
}

func (db *Database) GetIncompleteAudio(jobId int, speaker string) ([]*job.Job, error) {
	const q = `
	SELECT t.value, jt.no, jt.text_id
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

		if err := rows.Scan(&tmp.Content, &tmp.No, &tmp.TextId); err != nil {
			return nil, err
		}
		tmp.Speaker = speaker
		jobs = append(jobs, &tmp)
	}

	return jobs, nil
}