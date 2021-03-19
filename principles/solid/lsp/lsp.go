//
// Package srp shows how Liskov Substitution Principle can be applied in Go.
//
// https://en.wikipedia.org/wiki/Liskov_substitution_principle
//
package lsp

// Vehicle is interface describing vehicle that can be parked
type Vehicle interface {
	Park()
}

// Bicycle is simple vehicle that can be parked somewhere
type Bicycle struct {
	Vehicle
}

// Park bicycle
func (b Bicycle) Park() { /* ... */ }

// MountainBike is special bike designed for off-road
// While it function different from ordinary bicycle, it still have the same parking routine.
type MountainBike struct {
	Bicycle // embed Bicycle struct
}

// AdjustSuspension is mountain bike specific feature, let's say it is major difference
// of mountain bike from ordinary bicycle.
func (b Bicycle) AdjustSuspension() { /* ... */ }

// ----------------------------------------------------------------

// ParkingSpace for vehicles that accepts all vehicles of Vehicle interface
type ParkingSpace struct {
	parkedVehicles []Vehicle // list of currently parked vehicles
}

// ParkVehicle to parking lot
//
// This method is not concerned about particular type of vehicle, it only requires
// Vehicle interface with Park() method to be provided.
//
// We can pass object of type Bicycle but also of type MountainBike which is
// substitute of Bicycle which is main point of Liskov Substitution Principle.
func (s *ParkingSpace) ParkVehicle(v Vehicle) {
	// ...
	v.Park()
	// ...
	s.parkedVehicles = append(s.parkedVehicles, v)
}
