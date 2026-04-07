package db

import (
	"database/sql"
	"time"
	"APPOINMENT_BOOKING_SYSTEM/models"
)
func GetUserBookings(userID int) ([]models.BookingResponse, error) {

	rows, err := DB.Query(`
		SELECT b.id, b.coach_id, b.booking_datetime, c.name
		FROM bookings b
		JOIN coaches c ON b.coach_id = c.id
		WHERE b.user_id = ?
		ORDER BY b.booking_datetime ASC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.BookingResponse

	for rows.Next() {
		var b models.BookingResponse

		err := rows.Scan(&b.BookingID, &b.CoachID, &b.DateTime, &b.CoachName)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, b)
	}

	if bookings == nil {
		bookings = []models.BookingResponse{}
	}

	return bookings, nil
}


func InsertBooking(userID, coachID int, slot time.Time) error {
	_, err := DB.Exec(`
		INSERT INTO bookings (user_id, coach_id, booking_datetime)
		VALUES (?, ?, ?)
	`, userID, coachID, slot)

	return err
}

func InsertAvailability(id int, day string, start string, end string) error {
	_, err = DB.Exec(`
		INSERT INTO availability (coach_id, day_of_week, start_time, end_time)
		VALUES (?, ?, ?, ?)
	`, id, day, start, end)
	if err != nil {
		return err
	}
	return nil
}

func CheckCoachExists(id int) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM coaches WHERE id = ?", id).Scan(&count)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return false, nil
		// }
		return false, err
	}
	return count > 0, nil
}

func CheckAvailabilityExists(coachID int, day string) (bool, error) {
	var count int

	err := DB.QueryRow(`
		SELECT COUNT(*) 
		FROM availability 
		WHERE coach_id = ? AND day_of_week = ?
	`, coachID, day).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CheckMaxAvailabilityReached(coachID int) (bool, error) {
	var count int

	err := DB.QueryRow(`
		SELECT COUNT(*) 
		FROM availability 
		WHERE coach_id = ?
	`, coachID).Scan(&count)

	if err != nil {
		return false, err
	}

	return count >= 7, nil
}

func IsValidSlot(coachID int, slot time.Time) (bool, error) {
	day := slot.Weekday().String()

	var startStr, endStr string

	err := DB.QueryRow(`
		SELECT start_time, end_time
		FROM availability
		WHERE coach_id = ? AND day_of_week = ?
	`, coachID, day).Scan(&startStr, &endStr)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	start, _ := time.Parse("15:04:05", startStr)
	end, _ := time.Parse("15:04:05", endStr)

	// build full datetime
	startTime := time.Date(slot.Year(), slot.Month(), slot.Day(),
		start.Hour(), start.Minute(), 0, 0, time.UTC)

	endTime := time.Date(slot.Year(), slot.Month(), slot.Day(),
		end.Hour(), end.Minute(), 0, 0, time.UTC)

	// check slot inside range
	if slot.Before(startTime) || !slot.Before(endTime) {
		return false, nil
	}

	// check 30-min alignment
	if slot.Minute()%30 != 0 {
		return false, nil
	}

	return true, nil
}
