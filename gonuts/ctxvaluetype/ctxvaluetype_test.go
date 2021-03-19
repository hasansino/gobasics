package ctxvaluetype

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCtxValueType(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, strKey, "value")
	assert.Equal(t, "value", ctx.Value(strKey))
	assert.Nil(t, ctx.Value(typedKey)) // same value but different type

	ctx2 := context.Background()
	ctx2 = context.WithValue(ctx2, typedKey, "value")
	assert.Nil(t, ctx2.Value(strKey))
	assert.Equal(t, "value", ctx2.Value(typedKey))
}
