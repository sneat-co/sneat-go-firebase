package facade2firebase

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
)

func IssueFirebaseAuthToken(ctx context.Context, userID string, issuer string) (token string, err error) {

	var fbAuthClient *auth.Client
	if fbAuthClient, err = getFirebaseAuthClient(ctx); err != nil {
		err = fmt.Errorf("failed to get Firebase Auth client: %w", err)
		return
	}
	if issuer == "" {
		token, err = fbAuthClient.CustomToken(ctx, userID)
		if err != nil {
			err = fmt.Errorf("failed to create custom Firebase token for userID=%s without claims: %w", userID, err)
			return
		}
	} else {
		claims := map[string]any{}
		claims["issuer"] = issuer
		token, err = fbAuthClient.CustomTokenWithClaims(ctx, userID, claims)
		if err != nil {
			err = fmt.Errorf("failed to create custom Firebase token for userID=%s with claims (%+v): %w", userID, claims, err)
			return
		}
	}
	return
}
