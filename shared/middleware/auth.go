package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"messenger/shared/responsewriter"

	"github.com/golang-jwt/jwt/v5"
)

var ErrAuth = errors.New("authentication failed")

type Claims struct {
	Email string `json:"email"`
	Id    string `json:"id"`
	jwt.RegisteredClaims
}

const (
	UserIDKey = "id"
)

func ParseToken(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenRaw := r.Header.Get("Authorization")
			if tokenRaw == "" {
				tokenRaw = r.URL.Query().Get("token")
				if tokenRaw == "" {
					responsewriter.ErrorResponseWriter(w, ErrAuth, http.StatusUnauthorized)
					return
				}
				tokenRaw = "Bearer " + tokenRaw
			}

			tokenString := strings.TrimPrefix(tokenRaw, "Bearer ")
			if tokenString == tokenRaw {
				responsewriter.ErrorResponseWriter(w, ErrAuth, http.StatusUnauthorized)
				return
			}

			token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpected signing method")
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				responsewriter.ErrorResponseWriter(w, ErrAuth, http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(*Claims)
			if !ok {
				responsewriter.ErrorResponseWriter(w, ErrAuth, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.Id)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
