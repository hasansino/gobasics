package lsp

import "testing"

func TestLiskovSubstitutionPrinciple(t *testing.T) {
	var (
		b       = Bicycle{}
		mb      = MountainBike{}
		parking = ParkingSpace{make([]Vehicle, 0)}
	)

	parking.ParkVehicle(b)
	parking.ParkVehicle(mb)

	// let's try to create a very expensive and special mountain bike
	type verySpecialBike struct {
		MountainBike
	}

	// we still can park it easily, since verySpecialBike have all the same
	// properties as Bicycle
	parking.ParkVehicle(verySpecialBike{})
}
