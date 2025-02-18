package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	constant "github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/constants"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
)

const (
	showInterestQuery = `INSERT INTO interests (user_id,property_id, is_accepted, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5)`
	checkVacancyQuery       = `SELECT COUNT(*) FROM properties where property_id=$1 AND status='vacant'`
	interestedPropertyQuery = `SELECT p.property_id, p.lister_id, p.title, p.description, p.city, p.address, p.rent, p.facilities, p.images, 
	p.preferred_gender, p.status, p.vacancy FROM properties AS p INNER JOIN interests AS i ON p.property_id=i.property_id 
	INNER JOIN users AS u ON u.user_id=i.user_id WHERE u.user_id=$1`
	removeInterestQuery  = `DELETE FROM interests AS i USING users AS u WHERE u.user_id=$1 and i.property_id=$2`
	interestedUsersQuery = `SELECT u.name, u.phone, u.email, u.gender, u.city, u.role, u.required_vacancy, u.tags FROM users as u INNER JOIN interests as i on i.user_id=u.user_id INNER JOIN 
	properties as p ON p.property_id=i.property_id
	where p.property_id=$1`
	checkAccessQuery    = `select count(*) from properties where lister_id=$1 and property_id=$2`
	acceptInterestQuery = `UPDATE interests SET is_accepted=$3 WHERE property_id=$1 AND user_id=$2`
)

type interestRepo struct {
	DB *sql.DB
}

type InterestRepo interface {
	ShowInterest(ctx context.Context, propertyId int) error
	CheckVacancy(ctx context.Context, propertyId int) error
	GetInterestedProperties(ctx context.Context) ([]models.Property, error)
	RemoveInterest(ctx context.Context, propertyId int) error
	GetInterestedUsers(ctx context.Context, propertyId int) ([]models.User, error)
	CheckAccess(ctx context.Context, propertyId int) error
	AcceptInterest(ctx context.Context, userId int, propertyId int, status bool) error
}

func NewInterestRepo(db *sql.DB) InterestRepo {
	return &interestRepo{
		DB: db,
	}
}

// Allows User to add interest to a property
func (r *interestRepo) ShowInterest(ctx context.Context, propertyId int) error {
	err := r.CheckVacancy(ctx, propertyId)
	if err != nil {
		return errhandler.ErrInternalServer
	}

	userId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return errhandler.ErrInternalServer
	}

	var status bool = false
	_, err = r.DB.ExecContext(ctx, showInterestQuery, userId, propertyId, status, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

// Checks whether a property is vacant
func (r *interestRepo) CheckVacancy(ctx context.Context, propertyId int) error {
	var count = 0
	err := r.DB.QueryRowContext(ctx, checkVacancyQuery, propertyId).Scan(&count)
	if count == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}

	return nil
}

// Allows a user to retrieve all properties they added interest in
func (r *interestRepo) GetInterestedProperties(ctx context.Context) ([]models.Property, error) {
	userId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return nil, errhandler.ErrInternalServer
	}

	rows, err := r.DB.QueryContext(ctx, interestedPropertyQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []models.Property
	for rows.Next() {
		var property models.Property
		var propertyFacilities, propertyImages []byte

		err := rows.Scan(
			&property.PropertyId, &property.ListerId, &property.Title, &property.Description,
			&property.City, &property.Address, &property.Rent, &propertyFacilities, &propertyImages,
			&property.PreferredGender, &property.Status, &property.Vacancy,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(propertyFacilities, &property.Facilities)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(propertyImages, &property.Images)
		if err != nil {
			return nil, err
		}

		properties = append(properties, property)
	}

	return properties, nil
}

// Allows a user to remove interest from a property
func (r *interestRepo) RemoveInterest(ctx context.Context, propertyId int) error {
	userId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return errhandler.ErrInternalServer
	}

	_, err := r.DB.ExecContext(ctx, removeInterestQuery, userId, propertyId)
	if err != nil {
		return err
	}

	return nil
}

// Allows the lister to get profiles of all users who shown interest in their listed property
func (r *interestRepo) GetInterestedUsers(ctx context.Context, propertyId int) ([]models.User, error) {
	err := r.CheckAccess(ctx, propertyId)
	if err != nil {
		return nil, errhandler.ErrAuth
	}

	rows, err := r.DB.QueryContext(ctx, interestedUsersQuery, propertyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		var userTags []byte
		err = rows.Scan(&user.Name, &user.Phone, &user.Email, &user.Gender, &user.City, &user.Role, &user.RequiredVacancy, &userTags)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(userTags, &user.Tags)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Checks whether the lister has access to a property
func (r *interestRepo) CheckAccess(ctx context.Context, propertyId int) error {
	userId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return errhandler.ErrInternalServer
	}

	var count = 0
	err := r.DB.QueryRowContext(ctx, checkAccessQuery, userId, propertyId).Scan(&count)
	if count == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}

// Allows a user to accept or reject interest from a user
func (r *interestRepo) AcceptInterest(ctx context.Context, userId int, propertyId int, status bool) error {
	_, err := r.DB.ExecContext(ctx, acceptInterestQuery, propertyId, userId, status)
	if err != nil {
		return err
	}

	return nil
}
