package sneatfb

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/dal-go/dalgo/dal"
	"github.com/dal-go/dalgo2firestore"
	"github.com/sneat-co/sneat-go-core/apicore"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/sneatauth"
)

func InitFirebaseForSneat(projectID, dbName string) {
	if projectID == "" {
		panic("projectID is empty")
	}
	if dbName == "" {
		panic("dbName is empty")
	}
	apicore.GetAuthTokenFromHttpRequest = getSneatAuthTokenFromHttpRequest
	sneatauth.GetUserInfo = GetUserInfo

	facade.GetDatabase = func(ctx context.Context) (dal.DB, error) {
		if client, err := firestore.NewClient(ctx, projectID); err != nil {
			return nil, fmt.Errorf("failed to create Firestore client: %w", err)
		} else {
			return dalgo2firestore.NewDatabase(dbName, client), nil
		}
	}
}
