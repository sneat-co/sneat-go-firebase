package facade2firebase

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
)

func getFirebaseAuthClient(ctx context.Context) (*auth.Client, error) {
	if app, err := getFirebaseApp(ctx); err != nil {
		return nil, fmt.Errorf("faield to initializing Firebase app: %w", err)
	} else if client, err := app.Auth(ctx); err != nil {
		return nil, fmt.Errorf("faield to initializing Firebase Auth client: %w", err)
	} else {
		return client, nil
	}
}
