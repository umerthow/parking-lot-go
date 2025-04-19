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
	if err != nil {
		panic(err)
	}

	fmt.Println("hello from go")
}
