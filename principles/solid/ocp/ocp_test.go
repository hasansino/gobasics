package ocp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenClosedPrinciple(t *testing.T) {
	_v := _Vehicle{}
	_v.vType = "bike"
	assert.Equal(t, 2, _v.Wheels())
	assert.Equal(t, "bike", _v.Name())

	_v.vType = "car"
	assert.Equal(t, 4, _v.Wheels())
	assert.Equal(t, "car", _v.Name())

	// ---

	v := Vehicle{}
	assert.Equal(t, 0, v.Wheels())
	assert.Equal(t, "undefined", v.Name())

	b := Bike{}
	b.wheels = 2
	assert.Equal(t, 2, b.Wheels())
	assert.Equal(t, "bike", b.Name())

	c := Car{}
	c.wheels = 4
	assert.Equal(t, 4, c.Wheels())
	assert.Equal(t, "car", c.Name())
}
