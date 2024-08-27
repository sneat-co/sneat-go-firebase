package sneatfb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/sneat-co/sneat-go-core/facade"
)

var firebaseTokenContextKey = "firebaseToken"

// FirebaseTokenFromContext gets Firebase token from context
func FirebaseTokenFromContext(ctx context.Context) *auth.Token {
	v := ctx.Value(&firebaseTokenContextKey)
	if v == nil {
		return nil
	}
	return v.(*auth.Token)
}

func verifyIDToken(ctx context.Context, authClient *auth.Client, idToken string) (token *auth.Token, err error) {
	defer func() {
		if fail := recover(); fail != nil {
			err = fmt.Errorf("%w: %v", facade.ErrUnauthorized, fail)
		}
	}()
	token, err = authClient.VerifyIDToken(ctx, idToken)
	return
}

// NewContextWithFirebaseToken creates a new context with a Firebase token
func NewContextWithFirebaseToken(ctx context.Context, token *auth.Token) context.Context {
	return context.WithValue(ctx, &firebaseTokenContextKey, token)
}

// NewFirebaseAuthToken creates Firebase authentication token
var NewFirebaseAuthToken = newFirebaseAuthToken

type TokenStrProvider func() (tokenStr string, err error)

// NewFirebaseAuthToken creates a new Firebase Auth Token
func newFirebaseAuthToken(ctx context.Context, getTokenStr TokenStrProvider, authRequired bool) (token *auth.Token, err error) {
	if ctx == nil {
		panic("newFirebaseAuthToken(ctx=nil) - should be called with non nil context")
	}

	var tokenStr string
	if tokenStr, err = getTokenStr(); err != nil {
		return nil, fmt.Errorf("failed to get auth token string: %w", err)
	}

	if tokenStr == "" {
		if authRequired {
			return nil, fmt.Errorf("%w: authentication is required but request is missing Firebase tokenStr", facade.ErrUnauthorized)
		}
		return nil, nil
	}

	var fbApp *firebase.App
	if fbApp, err = firebase.NewApp(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to create *firebase.App: %w", err)
	}

	var fbAuthClient *auth.Client
	if fbAuthClient, err = fbApp.Auth(ctx); err != nil {
		return nil, fmt.Errorf("failed to create Firebase *auth.Client: %w", err)
	}

	//isDemoProject := strings.HasPrefix(googleCloudProjectID, "demo")
	if token, err = verifyIDToken(ctx, fbAuthClient, tokenStr); err != nil {
		if authRequired {
			return token, fmt.Errorf("failed to verify Firebase tokenStr: %v\ntokenStr: %s", err, tokenStr)
		}
	}

	if token == nil {
		if authRequired {
			return nil, errors.New("firebase.auth.Client.VerifyIDToken() returned nil error and nil token")
		}
	} else if token.UID == "" {
		if authRequired {
			s := new(bytes.Buffer)
			_ = json.NewEncoder(s).Encode(token)
			return nil, fmt.Errorf("no UserID, decoded Token: %s\n\n encoded tokenStr: %s", s.String(), tokenStr)
		}
	}
	return token, nil
}
