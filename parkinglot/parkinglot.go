package parkinglot

import (
	"errors"

	"github.com/umerthow/parking-lot-go/model"
)

type ParkingLot struct {
	floors  int
	rows    int
	columns int
	spots   [][][]model.ParkingSpot
}

func NewParkingLot(floors, rows, columns int) (*ParkingLot, error) {
	if floors < 1 || floors > 8 {
		return nil, errors.New("invalid number of floors (1-8)")
	}
	if rows < 1 || rows > 1000 {
		return nil, errors.New("invalid number of rows (1-1000)")
	}
	if columns < 1 || columns > 1000 {
		return nil, errors.New("invalid number of columns (1-1000)")
	}

	spots := make([][][]model.ParkingSpot, floors)

	return &ParkingLot{
		floors:  floors,
		rows:    rows,
		columns: columns,
		spots:   spots,
	}, nil
}
