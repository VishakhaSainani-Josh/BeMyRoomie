package validations

import (
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/go-playground/validator/v10"
)

func ValidateLoginRequestStruct(loginRequest models.LoginRequest) error {
	validate := validator.New()
	err := validate.Struct(loginRequest)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Email" {
				return errhandler.ErrInvalidEmail
			}
			if e.Field() == "Password" {
				return errhandler.ErrInvalidPassword
			}
		}
	}
	return nil
}

func ValidateRegisterUserStruct(user models.NewUserRequest) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Phone" {
				return errhandler.ErrInvalidPhone
			}
			if e.Field() == "Email" {
				return errhandler.ErrInvalidEmail
			}
			if e.Field() == "Password" {
				return errhandler.ErrInvalidPassword
			}
			if e.Field() == "Gender" {
				return errhandler.ErrInvalidGender
			}
			if e.Field() == "RequiredVacancy" {
				return errhandler.ErrInavlidVacancy
			} else {
				return errhandler.ErrRequired
			}
		}
	}
	return nil
}

func ValidatePreferenceStruct(preference models.NewPreferenceRequest) error {
	validate := validator.New()
	err := validate.Struct(preference)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "City" {
				return errhandler.ErrRequired
			}
		}
	}
	return nil
}

func ValidateUpdateProfileStruct(user models.User) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Phone" {
				return errhandler.ErrInvalidPhone
			}
			if e.Field() == "Email" {
				return errhandler.ErrInvalidEmail
			}
			if e.Field() == "Password" {
				return errhandler.ErrInvalidPassword
			}
			if e.Field() == "Gender" {
				return errhandler.ErrInvalidGender
			}
			if e.Field() == "RequiredVacancy" {
				return errhandler.ErrInavlidVacancy
			} else {
				return errhandler.ErrRequired
			}
		}
	}
	return nil
}

func ValidateRegisterPropertyStruct(property models.NewPropertyRequest) error {
	validate := validator.New()
	err := validate.Struct(property)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Title" {
				return errhandler.ErrInvalidTitle
			}
			if e.Field() == "Gender" {
				return errhandler.ErrInvalidGender
			}
			if e.Field() == "Vacancy" {
				return errhandler.ErrInavlidVacancy
			} else {
				return errhandler.ErrRequired
			}
		}
	}
	return nil
}

func ValidateUpdatePropertyStruct(property models.Property) error {
	validate := validator.New()
	err := validate.Struct(property)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			if e.Field() == "Title" {
				return errhandler.ErrInvalidTitle
			}
			if e.Field() == "Gender" {
				return errhandler.ErrInvalidGender
			}
			if e.Field() == "Vacancy" {
				return errhandler.ErrInavlidVacancy
			} else {
				return errhandler.ErrRequired
			}
		}
	}
	return nil
}

