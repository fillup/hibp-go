package hibp

import (
	"testing"

	"github.com/gofrs/uuid"
)

func Test_asHashes(t *testing.T) {
	tests := []struct {
		name       string
		pw         string
		wantPrefix string
		wantSuffix string
	}{
		{
			name: "pass123",
			pw: "pass123",
			wantPrefix: "aafdc",
			wantSuffix: "23870ecbcd3d557b6423a8982134e17927e",
		},
		{
			name: "alk4f2355f2d45d4f",
			pw: "alk4f2355f2d45d4f",
			wantPrefix: "2084a",
			wantSuffix: "389b33471898ad8ee44b47ed98ff2856053",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := asHashes(tt.pw)
			if got != tt.wantPrefix {
				t.Errorf("asHashes() got = %v, wantPrefix %v", got, tt.wantPrefix)
			}
			if got1 != tt.wantSuffix {
				t.Errorf("asHashes() got1 = %v, wantPrefix %v", got1, tt.wantSuffix)
			}
		})
	}
}

func TestIsPwned(t *testing.T) {

	validPw, _ := uuid.NewV4()

	tests := []struct {
		name    string
		pw    string
		want    bool
		wantErr bool
	}{
		{
			name: "pwned password",
			pw: "pass123",
			want: true,
			wantErr: false,
		},
		{
			name: "good password",
			pw: validPw.String(),
			want: false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsPwned(tt.pw)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsPwned() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsPwned() got = %v, want %v", got, tt.want)
			}
		})
	}
}