package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/constant"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/jwt"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.HandleResponse(w, http.StatusUnauthorized, "unauthorized access")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseJWT(tokenString)
		if err != nil {
			response.HandleResponse(w, http.StatusUnauthorized, "unauthorized access")
			return
		}

		userId := token.UserId
		ctx := context.WithValue(r.Context(), constant.UserIdKey, userId)
		role := token.Role
		ctx = context.WithValue(ctx, constant.RoleKey, role)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
