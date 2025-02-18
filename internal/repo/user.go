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
	selectQuery = `SELECT user_id, name, phone, email, password, gender,city,role,required_vacancy,tags FROM users 
	WHERE email = $1`
	updateQuery = `UPDATE users SET tags=$1,city=$2 
	where user_id=$3`
	viewProfileQuery = `SELECT name, phone, email, gender,city,role,required_vacancy,tags FROM users 
	WHERE user_id = $1`
	updateProfileQuery = `UPDATE users SET name=$2, phone=$3, email=$4, gender=$5, city=$6, role=$7, required_vacancy=$8,tags=$9 
	WHERE user_id=$1`
)

type userRepo struct {
	DB *sql.DB
}

type UserRepo interface {
	RegisterUser(ctx context.Context, user models.NewUserRequest) (int, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	AddPreferences(ctx context.Context, preference models.NewPreferenceRequest) error
	ViewProfile(ctx context.Context) (models.User, error)
	UpdateProfile(ctx context.Context, user models.User) error
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

func (r *userRepo) ViewProfile(ctx context.Context) (models.User, error) {
	var user models.User

	userIdValue := ctx.Value(constant.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		return user, errors.New("invalid or missing user ID in context")
	}
	var userTags []byte
	err := r.DB.QueryRowContext(ctx, viewProfileQuery, userId).Scan(&user.Name, &user.Phone, &user.Email, &user.Gender, &user.City,
		&user.Role, &user.RequiredVacancy, &userTags)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(userTags, &user.Tags)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepo) UpdateProfile(ctx context.Context, user models.User) error {
	existingProfile, err := r.ViewProfile(ctx)
	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = existingProfile.Name
	}
	if user.Phone == "" {
		user.Phone = existingProfile.Phone
	}
	if user.Email == "" {
		user.Email = existingProfile.Email
	}
	if user.Gender == "" {
		user.Gender = existingProfile.Gender
	}
	if user.City == "" {
		user.City = existingProfile.City
	}
	if user.Role == "" {
		user.Role = existingProfile.Role
	}
	if user.RequiredVacancy == 0 {
		user.RequiredVacancy = existingProfile.RequiredVacancy
	}
	if len(user.Tags) == 0 {
		user.Tags = existingProfile.Tags
	}

	userTags, err := json.Marshal(user.Tags)
	if err != nil {
		return err
	}

	userIdValue := ctx.Value(constant.UserIdKey)
	userId, ok := userIdValue.(int)
	if !ok {
		return errors.New("invalid or missing user ID in context")
	}

	_, err = r.DB.ExecContext(ctx, updateProfileQuery, userId, user.Name, user.Phone, user.Email, user.Gender, user.City, user.Role, user.RequiredVacancy, userTags)
	if err != nil {
		return err
	}

	return nil

}
