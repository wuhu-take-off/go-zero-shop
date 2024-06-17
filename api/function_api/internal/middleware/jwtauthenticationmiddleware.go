package middleware

import (
	"context"
	"net/http"
)

type JwtAuthenticationMiddleware struct {
}

func NewJwtAuthenticationMiddleware() *JwtAuthenticationMiddleware {
	return &JwtAuthenticationMiddleware{}
}

func (m *JwtAuthenticationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "token", authorization)

		// Passthrough to next handler if need
		next(w, r.WithContext(ctx))
	}
}
