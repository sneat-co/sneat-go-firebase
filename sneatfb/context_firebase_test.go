package sneatfb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSneatAuthTokenFromHttpRequest(t *testing.T) {
	t.Run("panics_on_nil_request", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _ = getSneatAuthTokenFromHttpRequest(nil, false)
		})
	})
}
