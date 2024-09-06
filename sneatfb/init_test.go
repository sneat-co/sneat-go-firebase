package sneatfb

import (
	"context"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitFirebaseForSneat(t *testing.T) {
	t.Run("panic_on_empty_project_id", func(t *testing.T) {
		assert.Panics(t, func() {
			InitFirebaseForSneat("", "dbName")
		})
	})
	t.Run("empty_dbName", func(t *testing.T) {
		InitFirebaseForSneat("projectID", "")
		db, err := facade.GetSneatDB(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, "default", db.ID())
	})
}
