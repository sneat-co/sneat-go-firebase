package sneatfb

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/sneat-co/sneat-go-core/sneatauth"
)

func GetUserInfo(ctx context.Context, uid string) (authUser *sneatauth.AuthUserInfo, err error) {
	fbApp, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return nil, err
	}
	auth, err := fbApp.Auth(ctx)
	if err != nil {
		return nil, err
	}
	fbUser, err := auth.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	authUser = &sneatauth.AuthUserInfo{
		AuthProviderUserInfo: &sneatauth.AuthProviderUserInfo{
			UID:         fbUser.UID,
			Email:       fbUser.Email,
			DisplayName: fbUser.DisplayName,
			PhoneNumber: fbUser.PhoneNumber,
			PhotoURL:    fbUser.PhotoURL,
		},
		EmailVerified:          fbUser.EmailVerified,
		Disabled:               fbUser.Disabled,
		CustomClaims:           fbUser.CustomClaims,
		TokensValidAfterMillis: fbUser.TokensValidAfterMillis,
	}
	if fbUser.UserMetadata != nil {
		authUser.UserMetadata = &sneatauth.UserMetadata{
			CreationTimestamp:    fbUser.UserMetadata.CreationTimestamp,
			LastLogInTimestamp:   fbUser.UserMetadata.LastLogInTimestamp,
			LastRefreshTimestamp: fbUser.UserMetadata.LastRefreshTimestamp,
		}
	}
	for _, pui := range fbUser.ProviderUserInfo {
		authUser.ProviderUserInfo = append(authUser.ProviderUserInfo, &sneatauth.AuthProviderUserInfo{
			UID:         pui.UID,
			Email:       pui.Email,
			DisplayName: pui.DisplayName,
			PhoneNumber: pui.PhoneNumber,
			PhotoURL:    pui.PhotoURL,
			ProviderID:  pui.ProviderID,
		})
	}
	if fbUser.MultiFactor != nil {
		authUser.MultiFactor = new(sneatauth.MultiFactorSettings)
		for _, ef := range fbUser.MultiFactor.EnrolledFactors {
			authUser.MultiFactor.EnrolledFactors = append(authUser.MultiFactor.EnrolledFactors, &sneatauth.MultiFactorInfo{
				UID:                 ef.UID,
				DisplayName:         ef.DisplayName,
				EnrollmentTimestamp: ef.EnrollmentTimestamp,
				FactorID:            ef.FactorID,
				PhoneNumber:         ef.PhoneNumber,
			})
		}
	}
	return
}
