package errhandler

import (
	"errors"
	"net/http"
)

var (
	Errbody            = errors.New("error reading body")
	ErrInvalidReq      = errors.New("error invalid request")
	ErrInternalServer  = errors.New("internal server error")
	ErrUserExist       = errors.New("email already registered")
	ErrUserInvalid     = errors.New("invalid credentials")
	ErrUserMissing     = errors.New("user doesn't exist")
	ErrHash            = errors.New("password hashing error")
	ErrToken           = errors.New("could not generate token")
	ErrAuth            = errors.New("unauthorized access")
	ErrExistProperty   = errors.New("already vacant property for user exists")
	ErrPropertyAccess  = errors.New("unauthorized to update property")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("password length should be between 8 to 20")
	ErrRequired        = errors.New("invalid details")
	ErrInvalidPhone    = errors.New("number length should be 10")
	ErrInvalidGender   = errors.New("gender should be Male or Female")
	ErrInavlidVacancy  = errors.New("required vacancy should be greater than 0")
	ErrInvalidTitle    = errors.New("title length should be atleast 2")
)

func MapError(err error) (int, string) {
	switch {
	case errors.Is(err, ErrInvalidReq):
		return http.StatusBadRequest, ErrInvalidReq.Error()
	case errors.Is(err, Errbody):
		return http.StatusInternalServerError, Errbody.Error()
	case errors.Is(err, ErrUserExist):
		return http.StatusConflict, ErrUserExist.Error()
	case errors.Is(err, ErrUserInvalid):
		return http.StatusUnauthorized, ErrUserInvalid.Error()
	case errors.Is(err, ErrUserMissing):
		return http.StatusNotFound, ErrUserMissing.Error()
	case errors.Is(err, ErrHash):
		return http.StatusInternalServerError, ErrHash.Error()
	case errors.Is(err, ErrToken):
		return http.StatusInternalServerError, ErrToken.Error()
	case errors.Is(err, ErrExistProperty):
		return http.StatusBadRequest, ErrExistProperty.Error()
	case errors.Is(err, ErrInternalServer):
		return http.StatusInternalServerError, ErrInternalServer.Error()
	case errors.Is(err, ErrPropertyAccess):
		return http.StatusBadRequest, ErrPropertyAccess.Error()
	case errors.Is(err, ErrAuth):
		return http.StatusUnauthorized, ErrAuth.Error()
	case errors.Is(err, ErrInvalidEmail):
		return http.StatusUnprocessableEntity, ErrInvalidEmail.Error()
	case errors.Is(err, ErrInvalidPassword):
		return http.StatusUnprocessableEntity, ErrInvalidPassword.Error()
	case errors.Is(err, ErrRequired):
		return http.StatusUnprocessableEntity, ErrRequired.Error()
	case errors.Is(err, ErrInvalidPhone):
		return http.StatusUnprocessableEntity, ErrInvalidPhone.Error()
	case errors.Is(err, ErrInvalidGender):
		return http.StatusUnprocessableEntity, ErrInvalidGender.Error()
	case errors.Is(err, ErrInavlidVacancy):
		return http.StatusUnprocessableEntity, ErrInavlidVacancy.Error()
	case errors.Is(err, ErrInvalidTitle):
		return http.StatusUnprocessableEntity, ErrInvalidTitle.Error()

	default:
		return http.StatusInternalServerError, "default internal server error"
	}
}
