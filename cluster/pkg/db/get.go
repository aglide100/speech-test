package db

import "github.com/aglide100/speech-test/cluster/pkg/request"

func (db *Database) GetParent(text string) (int, error) {
	const q = `
	SELECT id FROM Job WHERE text = $1
	`

	var parent int
	err := db.conn.QueryRow(q, text).Scan(&parent)
	if err != nil {
		return -1, err
	}	

	return parent, nil
}

func (db *Database) GetIncompleteJob() ([]request.Request, error) {
	const q = `
	SELECT j.id, j.speaker,
       GROUP_CONCAT(t.v SEPARATOR ' ') AS texts
	FROM Job AS j
	         LEFT JOIN (
	    SELECT parent, value as v
	    FROM Texts
	    ORDER BY no
	) AS t ON j.id = t.parent
	         LEFT JOIN (
	    SELECT parent, COUNT(*) AS audio_count
	    FROM Audio
	    GROUP BY parent
	) AS a ON j.id = a.parent
	WHERE j.max_index != a.audio_count OR a.audio_count IS NULL
	GROUP BY j.id`

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