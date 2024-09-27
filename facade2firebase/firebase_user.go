package facade2firebase

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/sneat-co/sneat-go-core/dto4auth"
)

func CreateFirebaseUser(ctx context.Context, userToCreate dto4auth.DataToCreateUser) (uid string, err error) {
	displayName := userToCreate.Names.FirstName
	if userToCreate.Names.LastName != "" {
		if displayName != "" {
			displayName += " "
		}
		displayName += userToCreate.Names.LastName
	}
	if displayName == "" {
		displayName = userToCreate.Names.UserName
	}
	firebaseUserToCreate := (&auth.UserToCreate{}).
		DisplayName(displayName)
	if userToCreate.PhotoURL != "" {
		firebaseUserToCreate = firebaseUserToCreate.PhotoURL(userToCreate.PhotoURL)
	}

	var fbAuthClient *auth.Client
	if fbAuthClient, err = getFirebaseAuthClient(ctx); err != nil {
		return
	}

	var firebaseUserRecord *auth.UserRecord
	if firebaseUserRecord, err = fbAuthClient.CreateUser(ctx, firebaseUserToCreate); err != nil {
		err = fmt.Errorf("failed to create firebase user: %w", err)
		return
	}
	customClaims := map[string]interface{}{
		"authProvider": userToCreate.AuthAccount.Provider,
	}
	if err = fbAuthClient.SetCustomUserClaims(ctx, firebaseUserRecord.UID, customClaims); err != nil {
		err = fmt.Errorf("failed to set custom claims: %w", err)
		return
	}
	uid = firebaseUserRecord.UID
	return
}

func DeleteFirebaseUser(ctx context.Context, uid string) (err error) {
	var fbAuthClient *auth.Client
	if fbAuthClient, err = getFirebaseAuthClient(ctx); err != nil {
		return
	}
	if err = fbAuthClient.DeleteUser(ctx, uid); err != nil {
		err = fmt.Errorf("failed to delete firebase user: %w", err)
	}
	return
}
