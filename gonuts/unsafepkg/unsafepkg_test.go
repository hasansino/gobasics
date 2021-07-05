package unsafepkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestB2i(t *testing.T) {
	assert.Equal(t, int8(0), B2i(false))
	assert.Equal(t, int8(1), B2i(true))
}

func BenchmarkB2i(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B2i(true)
		B2i(false)
	}
}

func BenchmarkB2iStandard(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_B2i(true)
		_B2i(false)
	}
}
