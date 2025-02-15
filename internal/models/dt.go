package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type NewUserRequest struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Gender          string `json:"gender"`
	Role            string `json:"role"`
	RequiredVacancy int    ` json:"required_vacancy"`
}

type UserResponse struct {
	UserId  int    `json:"userId"`
	Message string `json:"message"`
}

type NewPreferenceRequest struct {
	City string   `json:"city"`
	Tags []string `json:"tags"`
}

type User struct {
	UserId          int      `db:"user_id" json:"userId"`
	Name            string   `db:"name" json:"name"`
	Phone           string   `db:"phone" json:"phone"`
	Email           string   `db:"email" json:"email"`
	Password        string   `db:"passsword" json:"password"`
	Gender          string   `db:"gender" json:"gender"`
	City            string   `db:"city" json:"city"`
	Role            string   `db:"role" json:"role"`
	RequiredVacancy int      `db:"required_vacancy" json:"required_vacancy"`
	Tags            []string `db:"tags" json:"tags"`
}
