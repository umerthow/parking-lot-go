package parkinglot

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func (pl *ParkingLot) Park(vehicleType model.VehicleType, vehicleNumber string) (spotId string, err error) {
	for f := 0; f < pl.floors; f++ {
		for r := 0; r < pl.rows; r++ {
			for c := 0; c < pl.columns; c++ {
				spot := &pl.spots[f][r][c]
				if spot.Type == vehicleType && !spot.IsOccupied {
					spot.IsOccupied = true
					spot.VehicleNumber = vehicleNumber
					spot.OccupiedAt = time.Now()

					return spot.ID(), nil
				}
			}
		}
	}

	return "", errors.New("no available parking spot")
}

func (pl *ParkingLot) UnPark(spotId string, vehicleNumber string) error {
	floor, row, column, err := pl.parseSpotId(spotId)
	if err != nil {
		return err
	}

	spot := &pl.spots[floor-1][row-1][column-1]
	if !spot.IsOccupied || spot.VehicleNumber != vehicleNumber {
		return errors.New("vehicle not found at specified spot")
	}

	spot.IsOccupied = false
	spot.VehicleNumber = ""

	return nil
}

func (pl *ParkingLot) parseSpotId(spotID string) (floor, row, column int, err error) {
	parts := strings.Split(spotID, "-")
	if len(parts) != 3 {
		return 0, 0, 0, errors.New("invalid spot ID format")
	}

	floor, err1 := strconv.Atoi(parts[0])
	row, err2 := strconv.Atoi(parts[1])
	column, err3 := strconv.Atoi(parts[2])

	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, errors.New("invalid spot ID format")
	}

	return floor, row, column, nil
}

func (pl *ParkingLot) AvailableSpots(vehicleType model.VehicleType) []string {
	var available []string
	for f := 0; f < pl.floors; f++ {
		for r := 0; r < pl.rows; r++ {
			for c := 0; c < pl.columns; c++ {
				spot := pl.spots[f][r][c]
				if spot.Type == vehicleType && !spot.IsOccupied {
					available = append(available, spot.ID())
				}
			}
		}
	}
	return available
}

func (pl *ParkingLot) SearchParkVehicle(vehicleNumber string) (spot model.ParkingSpot, err error) {
	for f := 0; f < pl.floors; f++ {
		for r := 0; r < pl.rows; r++ {
			for c := 0; c < pl.columns; c++ {
				spot := pl.spots[f][r][c]
				if spot.VehicleNumber == vehicleNumber && spot.IsOccupied {
					return spot, nil
				}
			}
		}
	}

	return spot, errors.New("vehicle not found")
}
