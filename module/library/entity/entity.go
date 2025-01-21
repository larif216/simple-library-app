package entity

import "time"

type Book struct {
	Title         string
	Authors       []string
	EditionNumber string
}

type PickupSchedule struct {
	ID       int
	Book     Book
	DateTime time.Time
}

type CreatePickupScheduleRequest struct {
	EditionNumber string `json:"edition_number"`
	DateTime      string `json:"datetime"`
}

type CreatePickupScheduleResponse struct {
	Schedule *PickupSchedule
	Message  string
}
