package data

import (
	"context"
	"testing"
)

var ctx = context.Background()

func TestGetMetadata(t *testing.T) {
	t.Run("does not throw err", func(t *testing.T) {
		_, err := Get(ctx)
		if err != nil {
			t.Errorf("GetMetadata() failed: %v", err)
		}
	})
}
