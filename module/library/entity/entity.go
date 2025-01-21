package entity

import "time"

type Book struct {
	Title         string
	Authors       []string
	EditionNumber string
	IsAvailable   bool
}

type PickupSchedule struct {
	ID       int
	Book     Book
	DateTime time.Time
}
