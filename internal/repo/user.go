package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	constant "github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/constants"
)

const (
	registerUserQuery = `INSERT INTO users (name, phone, email, password, gender,role, required_vacancy,created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9)
	RETURNING user_id`
	selectQuery = `SELECT user_id, name, phone, email, password, gender,city,role,required_vacancy,tags FROM users WHERE email = $1`
	updateQuery = `UPDATE users SET tags=$1,city=$2 where user_id=$3`
)

type userRepo struct {
	DB *sql.DB
}

type UserRepo interface {
	RegisterUser(ctx context.Context, user models.NewUserRequest) (int, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	AddPreferences(ctx context.Context, preference models.NewPreferenceRequest) error
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		DB: db,
	}
}

// Registers User details in database
func (r *userRepo) RegisterUser(ctx context.Context, user models.NewUserRequest) (int, error) {
	var userId int

	err := r.DB.QueryRowContext(ctx, registerUserQuery, user.Name, user.Phone, user.Email, user.Password, user.Gender, user.Role, user.RequiredVacancy, time.Now(), time.Now()).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

// Retreive User details using email id
func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	var userTags []byte
	err := r.DB.QueryRowContext(ctx, selectQuery, email).Scan(&user.UserId, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Gender, &user.City, &user.Role, &user.RequiredVacancy, &userTags)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(userTags, &user.Tags)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Add User's preferences like city and tags
func (r *userRepo) AddPreferences(ctx context.Context, preference models.NewPreferenceRequest) error {
	userIdValue := ctx.Value(constant.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		return errors.New("invalid or missing user ID in context")
	}
	userTags, err := json.Marshal(preference.Tags)
	if err != nil {
		return err
	}
	_, err = r.DB.ExecContext(ctx, updateQuery, userTags, preference.City, userId)
	if err != nil {
		return err
	}
	return nil
}
