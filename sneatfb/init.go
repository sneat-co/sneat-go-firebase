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
	apicore.GetAuthTokenFromHttpRequest = getSneatAuthTokenFromHttpRequest
	sneatauth.GetUserInfo = GetUserInfo

	facade.GetSneatDB = func(ctx context.Context) (db dal.DB, err error) {
		var client *firestore.Client
		if dbName == "" {
			if client, err = firestore.NewClient(ctx, projectID); err != nil {
				return nil, fmt.Errorf("failed to create Firestore client for default DB: %w", err)
			}
		} else {
			if client, err = firestore.NewClient(ctx, projectID); err != nil {
				return nil, fmt.Errorf("failed to create Firestore client for %s database: %w", dbName, err)
			}
		}

		var dbID = dbName
		if dbID == "" {
			dbID = "default"
		}
		return dalgo2firestore.NewDatabase(dbID, client), nil
	}
}
