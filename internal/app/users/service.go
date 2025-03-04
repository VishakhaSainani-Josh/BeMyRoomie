package users

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/jwt"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
)

type Service interface {
	RegisterUser(ctx context.Context, user models.NewUserRequest, role string) (int, error)
	LoginUser(ctx context.Context, loginRequest models.LoginRequest) (string, models.User, error)
	AddPreferences(ctx context.Context, preferences models.NewPreferenceRequest) error
	ViewProfile(ctx context.Context) (models.User, error)
	UpdateProfile(ctx context.Context, user models.User) error
}

type service struct {
	userRepo repo.UserRepo
}

func NewService(userRepo repo.UserRepo) Service {
	return &service{userRepo: userRepo}
}

// Registers the User. First it checks whether user is already registered if not then it hashes the password and calls db layer to register the user
func (s *service) RegisterUser(ctx context.Context, user models.NewUserRequest, role string) (int, error) {
	_, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return 0, errhandler.ErrUserExist
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return 0, errhandler.ErrHash
	}
	user.Password = string(bytes)

	user.Role = role

	userId, err := s.userRepo.RegisterUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// Allows user to Signin. First it checks whether user is already registered, if yes then it compares the provided password with users registered password. It generates a JWT token for users session and returns it.
func (s *service) LoginUser(ctx context.Context, loginRequest models.LoginRequest) (string, models.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		return "", user, errhandler.ErrUserMissing
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", user, errhandler.ErrUserInvalid
	}

	token, err := jwt.GenerateJWT(user.UserId, user.Role)
	if err != nil {
		return "", user, errhandler.ErrToken
	}

	return token, user, nil
}

// Adds user preferences. Calls db layer to add preferences.
func (s *service) AddPreferences(ctx context.Context, preferences models.NewPreferenceRequest) error {
	err := s.userRepo.AddPreferences(ctx, preferences)
	if err != nil {
		return errhandler.ErrHash
	}
	return nil
}

func (s *service) ViewProfile(ctx context.Context) (models.User, error) {
	user, err := s.userRepo.ViewProfile(ctx)
	if err != nil {
		return user, errhandler.ErrInvalidReq
	}
	return user, nil
}

func (s *service) UpdateProfile(ctx context.Context, user models.User) error {
	err := s.userRepo.UpdateProfile(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
