# Parking Lot System

A concurrent parking lot manager written in Go, featuring thread-safe spot allocation, vehicle history tracking, and real-time availability checks.

## Features
- **Multi-floor parking** (configurable floors/rows/columns)
- **Thread-safe operations** (mutex-protected concurrent access)
- **Vehicle history tracking** (latest spot or full history)
- **Spot type configuration** (Bicycle, Motorcycle, Automobile)

## Installation
```bash
# Clone the repository
git clone https://github.com/umerthow/parking-lot-go.git
cd parking-lot-go

# Build (requires Go 1.20+)
go build
```

## Core Methods
| Method | Signature    | Description   |
| ---  | :--- | :--- |
| Park | Park(vehicleID string, spotType SpotType)   | Given a vehicle type, assign an empty parking spot id and map the vehicleNumber. spotId is floor-row-column. If no free spot is found, return an error. |
| Unpark | Unpark(spotID string)  | Removes vehicle from parking spot. Return an error for failure to unpark a vehicle.    |
| AvailableSpots | AvailableSpots(vehicleType model.VehicleType)  |  Display the free spots for each vehicle type.    |
| Search | Search(vehicleNumber string)  |  Search vehicle. If the vehicle has been unparked, get its last spotId    |


## Usage
``` go

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

	_ = pl.SetSpotType(1, 1, 1, model.Motorcycle)
	_ = pl.SetSpotType(1, 1, 2, model.Automobile)
	_ = pl.SetSpotType(1, 1, 3, model.Bicycle)


	// park
	spotId, err := pl.Park(model.Automobile, "B1284SS")
	if err != nil {
		panic(err)
	}
    
	spotVehicle, err := pl.SearchParkVehicle("B1284ST")
	if err != nil {
		fmt.Println("SearchParkVehicle failed:", err)
	}

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

	_, err = pl.Park(model.Automobile, "B1284SD")
	spotIdS, err := pl.Park(model.Automobile, "B1284SS")
	err = pl.UnPark(spotIdS, "B1284SS")

	lastSpot, err := pl.Search("B1284SS")
	if err != nil {
		fmt.Println("Search failed:", err)
	} else {
		fmt.Println("Last parked at:", lastSpot)
	}
}
```

## Running
```bash

# running go
go run main.go
```

