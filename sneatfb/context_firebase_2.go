package sneatfb

import (
	"errors"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/sneat-co/sneat-go-core/facade"
	"github.com/sneat-co/sneat-go-core/httpserver"
	"github.com/sneat-co/sneat-go-core/sneatauth"
	"net/http"
	"strings"
)

const authorizationHeaderName = "Authorization"
const bearerPrefix = "Bearer"

// getSneatAuthTokenFromHttpRequest creates a context with a Firebase ContactID token
var getSneatAuthTokenFromHttpRequest = func(r *http.Request) (token *sneatauth.Token, err error) {
	ctx := r.Context()
	if ctx == nil {
		return nil, errors.New("request returned nil context")
	}
	authHeader := r.Header.Get(authorizationHeaderName)
	if authHeader != "" {
		bearerToken, err := getBearerToken(authHeader)
		if err != nil {
			return nil, fmt.Errorf("failed to get bearer token from authorization header: %w", err)
		}
		var fbToken *auth.Token
		fbToken, err = NewFirebaseAuthToken(ctx, func() (string, error) {
			return bearerToken, nil
		}, false)
		if err != nil {
			return nil, fmt.Errorf("failed to get Firebase auth toke: %w", err)
		}
		token := sneatauth.Token{UID: fbToken.UID, Original: fbToken}
		sneatauth.NewContextWithAuthToken(ctx, &token)
	}
	return token, err
}

// newAuthContext creates new authentication context
//var newAuthContext = func(r *http.Request) (facade.AuthContext, error) {
//	fbIDToken := func() (string, error) {
//		return getBearerToken(r.Header.Get(authorizationHeaderName))
//	}
//	return NewFirebaseAuthContext(fbIDToken), nil
//}

func getBearerToken(authorizationHeader string) (token string, err error) {
	if authorizationHeader == "" {
		return "", facade.ErrNoAuthHeader
	}
	if !strings.HasPrefix(authorizationHeader, bearerPrefix) {
		return "", httpserver.ErrNotABearerToken
	}
	return authorizationHeader[len(bearerPrefix)+1:], nil
}
