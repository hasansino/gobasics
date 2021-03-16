package structalloc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestStructAlloc(t *testing.T) {
	assert.Equal(t, uintptr(24), unsafe.Sizeof(A{}))
	assert.Equal(t, uintptr(16), unsafe.Sizeof(B{}))
}
