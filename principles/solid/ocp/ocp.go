// Package ocp shows how Open-Closed Principle can be applied in Go.
//
// https://en.wikipedia.org/wiki/Open%E2%80%93closed_principle
package ocp

// _Vehicle is a struct describing a vehicle, it has field `vType` which says
// which type of vehicle it is.
// While this approach is straightforward and looks simple, everytime we need
// to add new type of vehicle, modification of internal methods of `_Vehicle`
// are required.
type _Vehicle struct {
	vType  string
	wheels int
}

// Wheels returns number of wheels vehicle have
func (v *_Vehicle) Wheels() int {
	switch v.vType {
	case "car":
		return 4
	case "bike":
		return 2
	default:
		return 0 // no type
	}
}

// Name of vehicle
func (v *_Vehicle) Name() string {
	return v.vType
}

// ----------------------------------------------------------------

// Vehicle is abstract struct describing vehicles with wheels.
// It has fields which is common to all vehicles, and can be described
// as abstract type for extension.
type Vehicle struct {
	wheels int
}

// Wheels returns number of wheels vehicle have
func (v *Vehicle) Wheels() int {
	return v.wheels
}

// Name returns specific vehicle name
func (v *Vehicle) Name() string {
	return "undefined"
}

// ----------------------------------------------------------------

// Bike is vehicle with 2 wheels.
// It extends Vehicle by overriding Name() method, while Vehicle structure
// remains unmodified.
type Bike struct {
	Vehicle
}

// Name is overridden method of Vehicle struct
func (b *Bike) Name() string {
	return "bike"
}

// ----------------------------------------------------------------

// Car is vehicle with 4 wheels
type Car struct {
	Vehicle
}

// Name is overridden method of Vehicle struct
func (c *Car) Name() string {
	return "car"
}

// Open is extension method for Vehicle
func (c *Car) Open() {}
