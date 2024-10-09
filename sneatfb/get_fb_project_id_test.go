package sneatfb

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestGetFirebaseProjectID(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		osEnviron []string
	}{
		{
			name:      "empty",
			want:      "demo-local-sneat-app",
			osEnviron: []string{},
		},
		{
			name:      "empty",
			want:      "abc123",
			osEnviron: []string{"FIREBASE_PROJECT_ID=abc123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, ev := range tt.osEnviron {
				vals := strings.Split(ev, "=")
				k, v := vals[0], vals[1]
				if err := os.Setenv(k, v); err != nil {
					t.Fatalf("failed to set env var %s=%s: %v", k, v, err)
				}
			}
			assert.Equalf(t, tt.want, GetFirebaseProjectID(), "GetFirebaseProjectID()")
		})
	}
}
