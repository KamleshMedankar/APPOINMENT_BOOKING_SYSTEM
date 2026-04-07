package models

import (
	"time"
)

type AppConfig struct {
	Ports                    []int
	MaxRequestsPerConnection int
	TotalConnections         int
	Server                   struct {
		Host string
		Port string
	}
	MySQL           Mysql
	RedisExpiration int
}
type Mysql struct{
	Address string
	Net string
	User string
	Password string
	DBName string
}
type Availability struct {
	CoachID   int    `json:"coach_id"`
	Day       string `json:"day_of_week"`
	Starttime string `json:"start_time"`
	Endtime   string `json:"end_time"`
}

type Booking struct {
	ID       int       `json:"booking_id"`
	UserID   int       `json:"user_id"`
	CoachID  int       `json:"coach_id"`
	DateTime time.Time `json:"datetime"`
}

type SlotResponse struct {
	Slots []time.Time `json:"slots"`
}

type BookingResponse struct {
	BookingID int       `json:"booking_id"`
	CoachID   int       `json:"coach_id"`
	CoachName string    `json:"coach_name"`
	DateTime  time.Time `json:"datetime"`
}