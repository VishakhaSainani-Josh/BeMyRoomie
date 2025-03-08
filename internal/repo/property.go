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
	registerPropertyQuery = `INSERT INTO properties (lister_id, title, description, city, address, rent, facilities, images, 
    preferred_gender, status, vacancy, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING property_id;`
	activePropertyQuery = `SELECT property_id, lister_id, title, description, city, address, rent, 
	facilities, images, preferred_gender, status, vacancy FROM properties 
	WHERE lister_id = $1 AND status = 'vacant'`
	updatePropertyQuery = `UPDATE properties SET title=$2, description=$3, city=$4, address=$5, rent=$6, facilities=$7, images=$8,
	preferred_gender=$9, status=$10, vacancy=$11
	WHERE property_id = $1`
	usersPropertiesQuery = `SELECT property_id, lister_id, title, description, city, address, rent, facilities, images, 
	preferred_gender, status, vacancy FROM properties 
	WHERE lister_id = $1`
	propertyAccessQuery = `SELECT COUNT(*) FROM properties WHERE property_id = $1 AND lister_id = $2`
	getPropertyQuery    = `SELECT property_id,lister_id,title,description,city,address,rent,facilities,images,preferred_gender,
	status,vacancy from properties 
	where property_id=$1`
	getAllPropertiesQuery = `SELECT property_id,lister_id,title,description,city,address,rent,facilities,images,preferred_gender,
	status,vacancy from properties 
	where status='vacant'`
)

type propertyRepo struct {
	DB *sql.DB
}

type PropertyRepo interface {
	RegisterProperty(ctx context.Context, property models.NewPropertyRequest) (int, error)
	GetActiveProperty(ctx context.Context) (models.Property, error)
	UpdateProperty(ctx context.Context, property models.Property) error
	CheckPropertyAccess(ctx context.Context, propertyId int) error
	GetAllProperties(ctx context.Context) ([]models.Property, error)
	GetParticularProperty(ctx context.Context, propertyId int) (models.Property, error)
	GetUsersProperties(ctx context.Context) ([]models.Property, error)
}

func NewPropertyRepo(db *sql.DB) PropertyRepo {
	return &propertyRepo{
		DB: db,
	}
}

// Registers Property details in database
func (r *propertyRepo) RegisterProperty(ctx context.Context, property models.NewPropertyRequest) (int, error) {
	var propertyId int
	listerId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return 0, errhandler.ErrInternalServer
	}

	propertyFacilities, err := json.Marshal(property.Facilities)
	if err != nil {
		return 0, err
	}

	propertyImages, err := json.Marshal(property.Images)
	if err != nil {
		return 0, err
	}

	err = r.DB.QueryRowContext(ctx, registerPropertyQuery, listerId, property.Title, property.Description, property.City,
		property.Address, property.Rent, propertyFacilities, propertyImages, property.PreferredGender, property.Status, property.Vacancy,
		time.Now(), time.Now()).Scan(&propertyId)
	if err != nil {
		return 0, err
	}

	return propertyId, nil
}

// Get Listers Active property (one with status as vacant)
func (r *propertyRepo) GetActiveProperty(ctx context.Context) (models.Property, error) {
	var property models.Property

	listerId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return property, errhandler.ErrInternalServer
	}

	var propertyFacilities []byte
	var propertyImages []byte

	err := r.DB.QueryRowContext(ctx, activePropertyQuery, listerId).Scan(&property.PropertyId, &property.ListerId, &property.Title,
		&property.Description, &property.City, &property.Address, &property.Rent, &propertyFacilities, &propertyImages,
		&property.PreferredGender, &property.Status, &property.Vacancy)
	if err != nil {
		return property, err
	}

	err = json.Unmarshal(propertyFacilities, &property.Facilities)
	if err != nil {
		return property, err
	}

	err = json.Unmarshal(propertyImages, &property.Images)
	if err != nil {
		return property, err
	}

	return property, nil
}

// First it retreives the properties existing details, then it assigns the exsiting details that are not updated and updates the details that are asked
func (r *propertyRepo) UpdateProperty(ctx context.Context, property models.Property) error {
	existingProperty, err := r.GetParticularProperty(ctx, property.PropertyId)
	if err != nil {
		return err
	}

	if property.Status == "vacant" {
		activeProperty, err := r.GetActiveProperty(ctx)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if activeProperty.PropertyId != 0 && activeProperty.PropertyId != property.PropertyId {
			return errhandler.ErrExistProperty
		}
	}

	if property.Title == "" {
		property.Title = existingProperty.Title
	}
	if property.Description == "" {
		property.Description = existingProperty.Description
	}
	if property.City == "" {
		property.City = existingProperty.City
	}
	if property.Address == "" {
		property.Address = existingProperty.Address
	}
	if property.Rent == 0 {
		property.Rent = existingProperty.Rent
	}
	if len(property.Facilities) == 0 {
		property.Facilities = existingProperty.Facilities
	}
	if len(property.Images) == 0 {
		property.Images = existingProperty.Images
	}
	if property.PreferredGender == "" {
		property.PreferredGender = existingProperty.PreferredGender
	}
	if property.Status == "" {
		property.Status = existingProperty.Status
	}
	if property.Vacancy == 0 {
		property.Vacancy = existingProperty.Vacancy
	}

	var propertyFacilities []byte
	var propertyImages []byte

	propertyFacilities, err = json.Marshal(property.Facilities)
	if err != nil {
		return err
	}

	propertyImages, err = json.Marshal(property.Images)
	if err != nil {
		return err
	}

	_, err = r.DB.ExecContext(ctx, updatePropertyQuery, property.PropertyId, property.Title, property.Description, property.City, property.Address,
		property.Rent, propertyFacilities, propertyImages, property.PreferredGender, property.Status, property.Vacancy)
	if err != nil {
		return err
	}

	return nil
}

// It checks whether a lister has access to update a property
func (r *propertyRepo) CheckPropertyAccess(ctx context.Context, propertyId int) error {
	listerId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return errhandler.ErrInternalServer
	}

	count := 0
	err := r.DB.QueryRowContext(ctx, propertyAccessQuery, propertyId, listerId).Scan(&count)
	if count == 0 {
		return errhandler.ErrPropertyAccess
	}
	if err != nil {
		return err
	}

	return nil
}

// Get all properties that are having vacancy
func (r *propertyRepo) GetAllProperties(ctx context.Context) ([]models.Property, error) {
	rows, err := r.DB.QueryContext(ctx, getAllPropertiesQuery)
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

// Get all properties posted by the lister
func (r *propertyRepo) GetUsersProperties(ctx context.Context) ([]models.Property, error) {
	listerId, ok := ctx.Value(constant.UserIdKey).(int)
	if !ok {
		return nil, errhandler.ErrInternalServer
	}

	rows, err := r.DB.QueryContext(ctx, usersPropertiesQuery, listerId)
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

// Get a particular property using the property id
func (r *propertyRepo) GetParticularProperty(ctx context.Context, propertyId int) (models.Property, error) {
	var existingProperty models.Property
	var propertyFacilities []byte
	var propertyImages []byte
	err := r.DB.QueryRowContext(ctx, getPropertyQuery, propertyId).Scan(&existingProperty.PropertyId, &existingProperty.ListerId,
		&existingProperty.Title, &existingProperty.Description, &existingProperty.City, &existingProperty.Address, &existingProperty.Rent,
		&propertyFacilities, &propertyImages, &existingProperty.PreferredGender, &existingProperty.Status, &existingProperty.Vacancy)
	if err != nil {
		return existingProperty, err
	}

	err = json.Unmarshal(propertyFacilities, &existingProperty.Facilities)
	if err != nil {
		return existingProperty, err
	}

	err = json.Unmarshal(propertyImages, &existingProperty.Images)
	if err != nil {
		return existingProperty, err
	}

	return existingProperty, nil
}
