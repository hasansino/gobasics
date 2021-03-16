package nilnotnil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNitNotNil(t *testing.T) {

	A := NewNilWorker("A")
	assert.NotEqual(t, nil, A)
	assert.Equal(t, (*WorkerA)(nil), A)
	// fmt.Printf("%v %T\n", A, A)

	B := NewNilWorker("B")
	assert.NotEqual(t, nil, B)
	assert.Equal(t, (*WorkerB)(nil), B)
	// fmt.Printf("%v %T\n", B, B)

	// C is of type nil without concrete type
	C := NewNilWorker("C")
	assert.Equal(t, nil, C)
	// fmt.Printf("%v %T\n", C, C)
}
