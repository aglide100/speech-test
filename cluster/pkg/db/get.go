package db

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