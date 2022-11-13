package middlewares

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ExtractAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, err := getUser(r)
		if err != nil && r.URL.Path != "/metrics" {
			logrus.WithError(err).Warn("failed to get username from token")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		logrus.Debugf("[%s] user in context: %s", r.Header.Get("X-Original-Remote-Addr"), username)
		// add username to context
		ctx := context.WithValue(r.Context(), "user", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUser(r *http.Request) (string, error) {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		return "", errors.Wrap(err, "getting token cookie")
	}

	token, _ := jwt.Parse(tokenCookie.Value, nil)
	//if err != nil {
	//	return "", errors.Wrap(err, "parsing token cookie")
	//}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to parse claims from token cookie")
	}

	call, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("failed to extract username from claims")
	}

	return call, nil
}
