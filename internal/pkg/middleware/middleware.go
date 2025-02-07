package middleware

import (
	"context"
	"net/http"
	"strings"

	constant "github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/constants"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/jwt"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

//Checks whether the user is valid, that is signed in before the user can access the next handler functionality
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
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
