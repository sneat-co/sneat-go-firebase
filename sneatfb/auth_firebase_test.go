package sneatfb

import (
	"testing"
)

func TestNewFirebaseAuthContext(t *testing.T) {
	type args struct {
		getFirebaseIDToken func() (string, error)
	}
	tests := []struct {
		name          string
		args          args
		expectToPanic bool
	}{
		{
			name: "ok",
			args: args{getFirebaseIDToken: func() (string, error) {
				return "token", nil
			}},
			expectToPanic: false,
		},
		{
			name:          "panics on empty token",
			args:          args{getFirebaseIDToken: nil},
			expectToPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectToPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("NewFirebaseAuthContext() did not panic")
					}
				}()
			}
			NewFirebaseAuthContext(tt.args.getFirebaseIDToken)
		})
	}
}
