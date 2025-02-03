package repo

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
)

type userRepo struct {
	DB *sql.DB
}

type UserRepo interface {
	RegisterUser(user models.User) (int, error)
	GetUserByEmail(email string) (models.User, error)
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
	var user models.User
	err := r.DB.QueryRow("SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
