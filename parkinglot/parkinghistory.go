package parkinglot

import (
	"github.com/umerthow/parking-lot-go/model"
)

type VehicleHistory struct {
	Data map[string][]model.ParkingSpot
}

func NewVehicleHistory() *VehicleHistory {
	return &VehicleHistory{
		Data: make(map[string][]model.ParkingSpot),
	}
}

func (vh *VehicleHistory) Record(vehicleNumber string, spot model.ParkingSpot) {
	vh.Data[vehicleNumber] = append(vh.Data[vehicleNumber], spot)
}

func (vh *VehicleHistory) Get(vehicleNumber string) ([]model.ParkingSpot, bool) {
	history, exists := vh.Data[vehicleNumber]
	return history, exists
}
