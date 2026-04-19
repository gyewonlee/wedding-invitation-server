package sqldb

import (
	"fmt"
	"time"
)

func initializeAttendanceTable() error {
	_, err := sqlDb.Exec(`
		CREATE TABLE IF NOT EXISTS attendance (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			side VARCHAR(10),
			name VARCHER(20),
			meal VARCHAR(20),
			count INTEGER,
			timestamp INTEGER
		)
	`)
	return err
}

func GetAllAttendance() ([]types.Attendance, error) {
	rows, err := sqlDb.Query(`
		SELECT id, side, name, meal, count, timestamp
		FROM attendance
		ORDER BY timestamp DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	attendances := []types.Attendance{}

	for rows.Next() {
		var a types.Attendance
		err := rows.Scan(&a.Id, &a.Side, &a.Name, &a.Meal, &a.Count, &a.Timestamp)
		if err != nil {
			return nil, err
		}
		attendances = append(attendances, a)
	}

	return attendances, nil
}

func CreateAttendance(side, name, meal string, count int) error {
	_, err := sqlDb.Exec(`
		INSERT INTO attendance (side, name, meal, count, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`, side, name, meal, count, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
