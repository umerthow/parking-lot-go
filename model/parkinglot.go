package model

import "time"

type VehicleType string

const (
	Bicycle    VehicleType = "B-1"
	Motorcycle VehicleType = "M-1"
	Automobile VehicleType = "A-1"
	Inactive   VehicleType = "X-0"
)

type ParkingSpot struct {
	Floor         int
	Row           int
	Column        int
	Type          VehicleType
	IsOccupied    bool
	VehicleNumber string
	OccupiedAt    time.Time
}
