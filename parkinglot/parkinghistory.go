package parkinglot

import (
	"sync"

	"github.com/umerthow/parking-lot-go/model"
)

type VehicleHistory struct {
	mu   sync.RWMutex
	Data map[string][]model.ParkingSpot
}

func NewVehicleHistory() *VehicleHistory {
	return &VehicleHistory{
		Data: make(map[string][]model.ParkingSpot),
	}
}

func (vh *VehicleHistory) Record(vehicleNumber string, spot model.ParkingSpot) {
	vh.mu.Lock()
	defer vh.mu.Unlock()

	vh.Data[vehicleNumber] = append(vh.Data[vehicleNumber], spot)
}

func (vh *VehicleHistory) Get(vehicleNumber string) ([]model.ParkingSpot, bool) {
	vh.mu.RLock()
	defer vh.mu.RUnlock()

	history, exists := vh.Data[vehicleNumber]
	return history, exists
}
