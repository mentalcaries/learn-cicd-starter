package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

  type test struct {
    name string
    input http.Header
    wantKey string
    wantErr bool
  }

  h1 := make(http.Header)
  h1.Set("Authorization", "ApiKey abc123")

  h2 := make(http.Header)
  h2.Set("Authorization", "")

  h3 := make(http.Header)
  h3.Set("NoAuthorization", "No token")

  tests := []test{
    { name: "Valid Api Key",input: h1, wantKey: "abc123", wantErr: false},
    { name: "Empty Api Key", input: h2, wantKey: "", wantErr: true},
    { name: "Authorization Header Absemt", input: h3, wantKey: "", wantErr: true},
  }

  for _, tc := range tests {
    got, err  := GetAPIKey(tc.input)
    if tc.wantErr {
      if err == nil {
        t.Fatalf("%s: expected error, got nil", tc.name)
      }
      return
    }
    if err != nil {
      t.Fatalf("%s: unexpected error: %v", tc.name, err)
    }
    
    if got != tc.wantKey{
      t.Fatalf("%s: got key : %q, expected: %q", tc.name, got, tc.wantKey)
    }
  }
}
