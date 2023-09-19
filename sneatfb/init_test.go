package sneatfb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitFirebaseForSneat(t *testing.T) {
	t.Run("panic_on_empty_project_id", func(t *testing.T) {
		assert.Panics(t, func() {
			InitFirebaseForSneat("", "dbName")
		})
	})
	t.Run("panic_on_empty_db_name", func(t *testing.T) {
		assert.Panics(t, func() {
			InitFirebaseForSneat("projectID", "")
		})
	})
}
