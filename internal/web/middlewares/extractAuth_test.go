package middlewares

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtractAuthMiddleware(t *testing.T) {
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name     string
		token    string
		want     string
		wantNext bool
	}{
		{
			name:     "happy path",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlBBMFdKRiIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJzZWNvbmRfZmFjdG9yX3NldCI6dHJ1ZSwic2Vjb25kX2ZhY3Rvcl9vayI6dHJ1ZSwic2Vjb25kX2ZhY3Rvcl9rZXkiOiIiLCJpc19hZG1pbiI6dHJ1ZSwiZXhwIjoxNjUzNDIwNDQ4LCJpYXQiOjE2NTA4Mjg0NDh9.9PHk5MJidgQnQKdUv4940Df7mQH8A23O5AP3HSXUJPQ",
			want:     "PA0WJF",
			wantNext: true,
		},
		{
			name:     "error path: no cookie",
			want:     "",
			wantNext: false,
		},
		{
			name:     "error path: no username in cookie",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
			want:     "",
			wantNext: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/", nil)
			if tt.token != "" {
				r.AddCookie(&http.Cookie{Name: "token", Value: tt.token})
			}

			nextCalled := false
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextCalled = true
				username := r.Context().Value("user").(string)
				assert.Equal(t, tt.want, username, "username from context")
			})

			ExtractAuthMiddleware(nextHandler).ServeHTTP(w, r)

			assert.Equal(t, tt.wantNext, nextCalled, "next handler called")
			if !tt.wantNext {
				assert.Equal(t, http.StatusForbidden, w.Result().StatusCode, "should return 403")
			}
		})
	}
}
