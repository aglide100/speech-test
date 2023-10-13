package db

import (
	"database/sql"

	"github.com/aglide100/speech-test/cluster/pkg/job"
)

func SaveAudio(db *sql.DB, audio *job.Request) error {
	const q = `
	INSERT INTO Audio () 
	`

	// _, err := db.Exec(q, audio.Text, )
	return nil
}

func SaveJob(db *sql.DB, req *job.Request) error {
	const q = `
	INSERT INTO Job (Job.text, Job.date) VALUES ($1, now())
	`

	_, err := db.Exec(q, req.Text)
	if err != nil {
		return err
	}

	return nil 
}