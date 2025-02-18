package models

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,min=8,max=20"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type NewUserRequest struct {
	Name            string `json:"name" validate:"required,min=1,max=100"`
	Phone           string `json:"phone" validate:"required,len=10"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8,max=20"`
	Gender          string `json:"gender" validate:"required,oneof=Male Female"  `
	Role            string `json:"role" `
	RequiredVacancy int    ` json:"required_vacancy" validate:"omitempty,gt=0"`
}

type UserResponse struct {
	UserId  int    `json:"userId"`
	Message string `json:"message"`
}

type NewPreferenceRequest struct {
	City string   `json:"city" validate:"required,alpha"`
	Tags []string `json:"tags"`
}

type User struct {
	UserId          int      `db:"user_id" json:"userId"`
	Name            string   `db:"name" json:"name" validate:"omitempty,min=2,max=100"`
	Phone           string   `db:"phone" json:"phone" validate:"omitempty,len=10" `
	Email           string   `db:"email" json:"email" validate:"omitempty,email"`
	Password        string   `db:"passsword" json:"password"`
	Gender          string   `db:"gender" json:"gender" validate:"omitempty,oneof=Male Female"`
	City            string   `db:"city" json:"city" validate:"omitempty,alpha"`
	Role            string   `db:"role" json:"role"`
	RequiredVacancy int      `db:"required_vacancy" json:"required_vacancy" validate:"omitempty,gt=0"`
	Tags            []string `db:"tags" json:"tags"`
}

type Property struct {
	PropertyId      int      `json:"property_id"`
	ListerId        int      `json:"lister_id"`
	Title           string   `json:"title" validate:"omitempty,min=2,max=100"`
	Description     string   `json:"description"`
	City            string   `json:"city" validate:"omitempty,alpha"`
	Address         string   `json:"address"`
	Rent            int64    `json:"rent"`
	Facilities      []string `json:"facilities"`
	Images          []string `json:"images"`
	PreferredGender string   `json:"preferred_gender" validate:"omitempty,oneof=Male Female"`
	Status          string   `json:"status"`
	Vacancy         int      `json:"vacancy" validate:"omitempty,gt=0"`
}

type NewPropertyRequest struct {
	PropertyId      int      `json:"property_id"`
	ListerId        int      `json:"lister_id"`
	Title           string   `json:"title" validate:"required,min=2,max=100"`
	Description     string   `json:"description"`
	City            string   `json:"city" validate:"required,alpha"`
	Address         string   `json:"address" validate:"required"`
	Rent            int64    `json:"rent" validate:"required"`
	Facilities      []string `json:"facilities"`
	Images          []string `json:"images" validate:"required"`
	PreferredGender string   `json:"preferred_gender" validate:"required,oneof=Male Female"`
	Status          string   `json:"status"`
	Vacancy         int      `json:"vacancy" validate:"required,gt=0"`
}

type PropertyResponse struct {
	PropertyId int    `json:"propertyId"`
	Message    string `json:"message"`
}

type InterestStatusRequest struct {
	IsAccepted bool `json:"is_accepted"`
}
