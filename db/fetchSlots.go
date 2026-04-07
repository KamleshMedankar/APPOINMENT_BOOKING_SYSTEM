package db

import (
	"log"
	"time"
)

func FetchAvailableSlots(coachID int, date string) ([]string, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	day := parsedDate.Weekday().String() // Monday, Tuesday...
	log.Println("Input date:", date)
	log.Println("Parsed day:", day)
	rows, err := DB.Query(`
		SELECT start_time, end_time
		FROM availability
		WHERE coach_id = ? AND day_of_week = ?
	`, coachID, day)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var slots []string

	for rows.Next() {
		var startStr, endStr string
		rows.Scan(&startStr, &endStr)

		start, _ := time.Parse("15:04:05", startStr)
		end, _ := time.Parse("15:04:05", endStr)

		current := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
			start.Hour(), start.Minute(), 0, 0, time.UTC)

		endTime := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(),
			end.Hour(), end.Minute(), 0, 0, time.UTC)

		for current.Before(endTime) {
			slots = append(slots, current.Format(time.RFC3339))
			current = current.Add(30 * time.Minute)
		}
	}

	return slots, nil
}
