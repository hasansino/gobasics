package bitmask

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitmask(t *testing.T) {
	// Let's print values in decimal notation
	// fmt.Printf("%v\n", Capabilities)
	// [1 2 4 8 16 32 64]
	// As we can see all numbers except 1 is product of 2^x

	// Now print all values in binary format
	// for _, v := range Capabilities {
	//	fmt.Printf("%b\n", v)
	// }

	// 1
	// 10
	// 100
	// 1000
	// 10000
	// 100000
	// 1000000

	// Let's define an empty mask
	// Since it is empty (all bits are zeroed), we assume all capabilities are disabled.
	var mask uint64

	// Add A capability to mask
	mask = AddCapability(mask, CapabilityA) // 0 | 1 == 1
	assert.Equal(t, true, IsCapable(mask, CapabilityA))

	// Add more capabilities
	mask = AddCapability(mask, CapabilityB|CapabilityC|CapabilityD)
	assert.Equal(t, true, IsCapable(mask, CapabilityB))
	assert.Equal(t, true, IsCapable(mask, CapabilityC))
	assert.Equal(t, true, IsCapable(mask, CapabilityD))

	// Adding CapabilityA again will not change anything
	mask = AddCapability(mask, CapabilityA)
	assert.Equal(t, true, IsCapable(mask, CapabilityA))

	// Remove capability
	mask = RemoveCapability(mask, CapabilityA)
	assert.Equal(t, false, IsCapable(mask, CapabilityA))
}
