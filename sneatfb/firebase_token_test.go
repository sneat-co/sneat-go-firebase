package sneatfb

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"testing"
)

func TestFirebaseTokenFromContext(t *testing.T) {
	token := &auth.Token{}
	ctx := context.WithValue(context.Background(), &firebaseTokenContextKey, token)

	actual := FirebaseTokenFromContext(ctx)
	if actual != token {
		t.Errorf("actual token is not equal to expected")
	}
}
