package middleware

import (
	"context"
	"net/http"
	"strings"

	constant "github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/constants"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/jwt"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/response"
)

// Checks whether the user is valid, that is signed in before the user can access the next handler functionality
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.HandleResponse(w, http.StatusUnauthorized, errhandler.ErrAuth.Error(), nil)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseJWT(tokenString)
		if err != nil {
			response.HandleResponse(w, http.StatusUnauthorized, errhandler.ErrAuth.Error(), nil)
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

// Checks whether the user is authorized to perform an action, if role of user is lister then it allows property methods
func AuthorizeLister(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		role, ok := ctx.Value(constant.RoleKey).(string)
		if !ok {
			response.HandleResponse(w, http.StatusUnauthorized, errhandler.ErrInternalServer.Error(), nil)
			return
		}

		if role != "lister" {
			response.HandleResponse(w, http.StatusUnauthorized, errhandler.ErrAuth.Error(), nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthorizeFinder(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		role, ok := ctx.Value(constant.RoleKey).(string)
		if !ok {
			response.HandleResponse(w, http.StatusUnauthorized, errhandler.ErrInternalServer.Error(), nil)
			return
		}

		if role != "finder" {
			response.HandleResponse(w, http.StatusUnauthorized, errhandler.ErrAuth.Error(), nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func EnableCORS(next *http.ServeMux) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
