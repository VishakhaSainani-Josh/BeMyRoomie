package properties

import (
	"context"
	"database/sql"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/pkg/errhandler"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
)

type Service interface {
	RegisterProperty(ctx context.Context, property models.NewPropertyRequest) (int, error)
	UpdateProperty(ctx context.Context, property models.Property, propertyId int) error
	GetAllProperties(ctx context.Context) ([]models.Property, error)
	GetUsersProperties(ctx context.Context) ([]models.Property, error)
	GetParticularProperty(ctx context.Context, propertyId int) (models.Property, error)
}

type service struct {
	propertyRepo repo.PropertyRepo
}

func NewService(propertyRepo repo.PropertyRepo) Service {
	return &service{propertyRepo: propertyRepo}
}

/*
It first check that whether a lister is not having any active property(one with status as vacant) if not then it allows
the lister to register a new property
*/
func (s *service) RegisterProperty(ctx context.Context, property models.NewPropertyRequest) (int, error) {
	activeProperty, err := s.propertyRepo.GetActiveProperty(ctx)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if activeProperty.PropertyId != 0 {
		return 0, errhandler.ErrExistProperty
	}

	propertyId, err := s.propertyRepo.RegisterProperty(ctx, property)
	if err != nil {
		return 0, err
	}
	return propertyId, nil
}

// Allows a lister to update details of their posted property
func (s *service) UpdateProperty(ctx context.Context, property models.Property, propertyId int) error {
	err := s.propertyRepo.CheckPropertyAccess(ctx, propertyId)
	if err != nil {
		return err
	}

	property.PropertyId = propertyId
	err = s.propertyRepo.UpdateProperty(ctx, property)
	if err != nil {
		return err
	}
	return nil
}

// This Service is used to get all properties with vacancy
func (s *service) GetAllProperties(ctx context.Context) ([]models.Property, error) {
	properties, err := s.propertyRepo.GetAllProperties(ctx)
	if err != nil {
		return properties, err
	}
	return properties, nil
}

// This Service allows a lister to get all properties posted by them
func (s *service) GetUsersProperties(ctx context.Context) ([]models.Property, error) {
	properties, err := s.propertyRepo.GetUsersProperties(ctx)
	if err != nil {
		return properties, err
	}
	return properties, nil
}

// Allows to retreive a particular property with a propertyId
func (s *service) GetParticularProperty(ctx context.Context, propertyId int) (models.Property, error) {
	property, err := s.propertyRepo.GetParticularProperty(ctx, propertyId)
	if err != nil {
		return property, err
	}
	return property, nil
}
