package bitmask

// https://en.wikipedia.org/wiki/Mask_(computing)

// Let's represent some abstract capabilities
// Each of them will represent a specific bit
const (
	CapabilityA = 1 << iota
	CapabilityB
	CapabilityC
	CapabilityD
)

// IsCapable performs AND operation on m and c
// The result is true if at least one of c bits is set in m
//
// Example:
// m (01001) & c (1) => true
// m (00001) & c (01) => false
func IsCapable(m, c uint64) bool {
	return m&c != 0
}

// AddCapability performs binary OR operation between m and c
// Basically it sets c bits in m
//
// Example:
// m (1000000) | c (1001100) => 1001100
func AddCapability(m, c uint64) uint64 {
	return m | c
}

// RemoveCapability performs AND NOT operation between m and c.
// This zeroes bits set in c from m.
//
// Example:
// m (011011) &^ c (010010) => 001001
func RemoveCapability(m, c uint64) uint64 {
	return m &^ c
}
