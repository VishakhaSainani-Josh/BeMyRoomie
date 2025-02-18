package interests

import (
	"context"

	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/models"
	"github.com/VishakhaSainani-Josh/BeMyRoomie/internal/repo"
)

type Service interface {
	ShowInterest(ctx context.Context, propertyId int) error
	GetInterestedProperties(ctx context.Context) ([]models.Property, error)
	RemoveInterest(ctx context.Context, propertyId int) error
	GetInterestedUsers(ctx context.Context, propertyId int) ([]models.User, error)
	AcceptInterest(ctx context.Context, userId int, propertyId int, status bool) error
}

type service struct {
	interestRepo repo.InterestRepo
}

func NewService(interestRepo repo.InterestRepo) Service {
	return &service{interestRepo: interestRepo}
}

// This service allows a user to show interest in a property
func (s *service) ShowInterest(ctx context.Context, propertyId int) error {
	err := s.interestRepo.ShowInterest(ctx, propertyId)
	if err != nil {
		return err
	}
	return nil
}

// This service fetches the properties a user has shown interest in
func (s *service) GetInterestedProperties(ctx context.Context) ([]models.Property, error) {
	properties, err := s.interestRepo.GetInterestedProperties(ctx)
	if err != nil {
		return nil, err
	}
	return properties, nil
}

// This service allows a user to remove interest from a property
func (s *service) RemoveInterest(ctx context.Context, propertyId int) error {
	err := s.interestRepo.RemoveInterest(ctx, propertyId)
	if err != nil {
		return err
	}
	return nil
}

// This service allows a lister to get users shown interest in their listed property
func (s *service) GetInterestedUsers(ctx context.Context, propertyId int) ([]models.User, error) {
	users, err := s.interestRepo.GetInterestedUsers(ctx, propertyId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// This service allows lister to accept or reject interest from a user
func (s *service) AcceptInterest(ctx context.Context, userId int, propertyId int, status bool) error {
	err := s.interestRepo.AcceptInterest(ctx, userId, propertyId, status)
	if err != nil {
		return err
	}
	return nil
}
