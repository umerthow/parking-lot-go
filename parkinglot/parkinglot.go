package parkinglot

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/umerthow/parking-lot-go/model"
)

type ParkingLot struct {
	floors  int
	rows    int
	columns int
	spots   [][][]model.ParkingSpot
}

// Init Parking Lot
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

	for f := 0; f < floors; f++ {
		spots[f] = make([][]model.ParkingSpot, rows)
		for r := 0; r < rows; r++ {
			spots[f][r] = make([]model.ParkingSpot, columns)
			for c := 0; c < columns; c++ {
				// Default to inactive spot
				spots[f][r][c] = model.ParkingSpot{
					Floor:      f + 1,
					Row:        r + 1,
					Column:     c + 1,
					Type:       model.Inactive,
					IsOccupied: false,
				}
			}
		}
	}

	return &ParkingLot{
		floors:  floors,
		rows:    rows,
		columns: columns,
		spots:   spots,
	}, nil
}

func (pl *ParkingLot) SetSpotType(floor, row, column int, spotType model.VehicleType) error {
	if floor < 1 || floor > pl.floors {
		return errors.New("invalid floor number")
	}
	if row < 1 || row > pl.rows {
		return errors.New("invalid row number")
	}
	if column < 1 || column > pl.columns {
		return errors.New("invalid column number")
	}

	pl.spots[floor-1][row-1][column-1].Type = spotType

	pls, _ := json.MarshalIndent(pl.spots, "", "/")
	fmt.Println("spot", string(pls))
	return nil
}

func (pl *ParkingLot) Park(vehicleType model.VehicleType, vehicleNumber string) (string, error) {
	for f := 0; f < pl.floors; f++ {
		for r := 0; r < pl.rows; r++ {
			for c := 0; c < pl.columns; c++ {
				spot := &pl.spots[f][r][c]
				if spot.Type == vehicleType && !spot.IsOccupied {
					spot.IsOccupied = true
					spot.VehicleNumber = vehicleNumber
					spot.OccupiedAt = time.Now()

					return fmt.Sprintf("%d-%d-%d", spot.Floor, spot.Row, spot.Column), nil
				}
			}
		}
	}

	return "", errors.New("no available parking spot")
}
