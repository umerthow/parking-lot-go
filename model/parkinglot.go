package model

import (
	"fmt"
	"time"
)

type VehicleType string

const (
	Bicycle    VehicleType = "B-1"
	Motorcycle VehicleType = "M-1"
	Automobile VehicleType = "A-1"
	Inactive   VehicleType = "0-0"
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

func (ps ParkingSpot) ID() string {
	return formatSpotID(ps.Floor, ps.Row, ps.Column)
}

func formatSpotID(floor, row, column int) string {
	return fmt.Sprintf("%d-%d-%d", floor, row, column)
}
