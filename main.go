package main

import (
	"fmt"

	"github.com/umerthow/parking-lot-go/model"
	"github.com/umerthow/parking-lot-go/parkinglot"
)

func main() {

	pl, err := parkinglot.NewParkingLot(1, 2, 10)

	if err != nil {
		panic(err)
	}

	err = pl.SetSpotType(1, 1, 1, model.Motorcycle)
	err = pl.SetSpotType(1, 1, 2, model.Automobile)
	err = pl.SetSpotType(1, 1, 3, model.Bicycle)
	err = pl.SetSpotType(1, 1, 4, model.Automobile)
	err = pl.SetSpotType(1, 2, 1, model.Automobile)
	if err != nil {
		panic(err)
	}

	// park
	spotId, err := pl.Park(model.Automobile, "B1284SS")
	if err != nil {
		panic(err)
	}
	fmt.Println("Spot Parked Occupied at %s", spotId)

	// get available spots
	availableSpots := pl.AvailableSpots(model.Automobile)
	fmt.Println("Available Spots - ", availableSpots)

	// unpark
	err = pl.UnPark(spotId, "B1284SS")
	if err != nil {
		fmt.Println("Unparking failed:", err)
		panic(err)
	}
	fmt.Println("Unparking success:", spotId)

	fmt.Println()
	fmt.Println("hello from go")
}
