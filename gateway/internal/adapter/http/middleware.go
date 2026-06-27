package http

import (
	"context"
	"net/http"
	"strings"
	authpb "task-tracker/gen/proto/auth"
)

func AuthMiddleware(authClient authpb.AuthServiceClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			token := strings.TrimPrefix(header, "Bearer ")

			if token == "" {
				writeError(w, http.StatusUnauthorized, "token missed")
				return
			}

			resp, err := authClient.ValidateToken(r.Context(), &authpb.ValidateTokenRequest{AccessToken: token})
			if err != nil {
				writeError(w, http.StatusUnauthorized, "invalid token")
				return
			}

			ctx := context.WithValue(r.Context(), "userID", resp.UserId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
