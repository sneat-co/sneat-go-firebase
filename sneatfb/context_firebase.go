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
func getSneatAuthTokenFromHttpRequest(r *http.Request, authRequired bool) (token *sneatauth.Token, err error) {
	if r == nil {
		panic("request is nil")
	}
	ctx := r.Context()
	if ctx == nil {
		return nil, errors.New("request returned nil context")
	}
	if authHeader := r.Header.Get(authorizationHeaderName); authHeader != "" {
		var bearerToken string
		if bearerToken, err = getBearerToken(authHeader); err != nil {
			return nil, fmt.Errorf("failed to get bearer token from authorization header: %w", err)
		}
		var fbToken *auth.Token

		getTokenStr := func() (string, error) {
			return bearerToken, nil
		}
		fbToken, err = NewFirebaseAuthToken(ctx, getTokenStr, authRequired)
		if err != nil {
			return nil, fmt.Errorf("failed to get Firebase auth toke: %w", err)
		}

		if fbToken != nil {
			token = &sneatauth.Token{
				UID:      fbToken.UID,
				Original: fbToken,
			}
		}
	}
	return
}

// newAuthContext creates new authentication context
//var newAuthContext = func(r *http.Request) (facade.AuthContext, error) {
//	getFirebaseIDToken := func() (string, error) {
//		return getBearerToken(r.Header.Get(authorizationHeaderName))
//	}
//	return NewFirebaseAuthContext(getFirebaseIDToken), nil
//}

func getBearerToken(authorizationHeader string) (token string, err error) {
	if authorizationHeader == "" {
		return "", facade.ErrNoAuthHeader
	}
	if !strings.HasPrefix(authorizationHeader, bearerPrefix) {
		return "", httpserver.ErrNotABearerToken
	}
	token = authorizationHeader[len(bearerPrefix)+1:]
	return strings.TrimSpace(token), nil
}
