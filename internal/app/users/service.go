package users

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/jwt"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
)

const (
	userExist   = "email already registered"
	userInvalid = "invalid crredentials"
	userMissing = "user doesn't exist"
	hashError   = "password hashing error"
	tokenErr    = "could not generate token"
)

type Service interface {
	RegisterUser(user models.User, role string) (int, error)
	LoginUser(email, password string) (string, error)
}

type service struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) Service {
	return &service{userRepo: userRepo}
}

func (s *service) RegisterUser(user models.User, role string) (int, error) {
	_, err := s.userRepo.GetUserByEmail(user.Email)
	if err == nil {
		return 0, errhandler.CustomError{
			Message:    userExist,
			StatusCode: http.StatusConflict,
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return 0, errhandler.CustomError{
			Message:    hashError,
			StatusCode: http.StatusInternalServerError,
		}
	}
	user.Password = string(bytes)

	user.Role = role

	userId, err := s.userRepo.RegisterUser(user)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (s *service) LoginUser(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errhandler.CustomError{
			Message:    userMissing,
			StatusCode: http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errhandler.CustomError{
			Message:    userInvalid,
			StatusCode: http.StatusUnauthorized,
		}
	}

	token, err := jwt.GenerateJWT(user.Email)
	if err != nil {
		return "", errhandler.CustomError{
			Message:    tokenErr,
			StatusCode: http.StatusInternalServerError,
		}
	}

	return token, nil
}
