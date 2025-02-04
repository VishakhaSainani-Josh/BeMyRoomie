package repo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
)

type userRepo struct {
	DB *sql.DB
}

type UserRepo interface {
	RegisterUser(user models.User) (int, error)
	GetUserByEmail(email string) (models.User, error)
	AddPreferences(userId int, tags []string, city string) error
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		DB: db,
	}
}

func (r *userRepo) RegisterUser(user models.User) (int, error) {
	registerUserQuery := `INSERT INTO users (name, phone, email, password, gender, city, role, required_vacancy, tags, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING user_id`

	var userId int

	userTags, err := json.Marshal(user.Tags)
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRow(registerUserQuery, user.Name, user.Phone, user.Email, user.Password, user.Gender, user.City,
		user.Role, user.RequiredVacancy, userTags, time.Now(), time.Now()).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *userRepo) GetUserByEmail(email string) (models.User, error) {
	selectQuery := `SELECT user_id, name, phone, email, password, gender, city, role, required_vacancy, tags FROM users WHERE email = $1`

	var user models.User
	var userTags []byte
	err := r.DB.QueryRow(selectQuery, email).Scan(&user.UserId, &user.Name, &user.Phone, &user.Email, &user.Password, &user.Gender, &user.City, &user.Role, &user.RequiredVacancy, &userTags)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(userTags, &user.Tags)
	if err != nil {
		fmt.Println("JSON Unmarshal Error:", err)
		return user, err
	}
	return user, nil
}

func (r *userRepo) AddPreferences(userId int, tags []string, city string) error {
	userTags, err := json.Marshal(tags)
	if err != nil {
		return err
	}

	updateQuery := `UPDATE users SET tags=$1,city=$2 where user_id=$3`
	_, err = r.DB.Exec(updateQuery, userTags, city, userId)
	if err != nil {
		return err
	}
	return nil
}
