package sneatfb

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"github.com/sneat-co/sneat-go-core/facade"
)

type firebaseAuthContext struct {
	getFirebaseIDToken func() (string, error)
	err                error
	fbAuthToken        *auth.Token
	userCtx            facade.UserContext
}

var _ facade.AuthContext = (*firebaseAuthContext)(nil)

func (v *firebaseAuthContext) User(ctx context.Context, authRequired bool) (userCtx facade.UserContext, err error) {
	if v.userCtx != nil || v.err != nil {
		return v.userCtx, v.err
	}
	var fbAuthToken *auth.Token
	if fbAuthToken, err = v.getFbAuthToken(ctx, authRequired); err != nil {
		return v.userCtx, v.err
	}
	v.userCtx = facade.NewContextWithUser(ctx, fbAuthToken.UID).User()
	return v.userCtx, nil
}

func (v *firebaseAuthContext) getFbAuthToken(ctx context.Context, authRequired bool) (*auth.Token, error) {
	if v.fbAuthToken == nil {
		v.fbAuthToken, v.err = NewFirebaseAuthToken(ctx, v.getFirebaseIDToken, authRequired)
	}
	return v.fbAuthToken, v.err
}

// NewFirebaseAuthContext creates new context with getFirebaseIDToken
func NewFirebaseAuthContext(getFirebaseIDToken func() (string, error)) facade.AuthContext {
	if getFirebaseIDToken == nil {
		panic("getFirebaseIDToken is nil")
	}
	return &firebaseAuthContext{getFirebaseIDToken: getFirebaseIDToken}
}
