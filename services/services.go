package services

import (
	"context"
	"fmt"

	"test-school/models"
	"test-school/repository"
)

type service struct {
	db repository.Repository
}

type Services interface {
	GetProfile(ctx context.Context, userID int) (*models.Profile, error)
	GetProfileList(ctx context.Context) ([]models.Profile, error)

	GetAuthKey(ctx context.Context, authKey string) (bool, error)
}

func NewService(db repository.Repository) Services {
	return service{db: db}
}

func (s service) GetProfile(ctx context.Context, userID int) (*models.Profile, error) {
	profile, err := s.db.GetProfileByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get profile error: %w", err)
	}

	return profile, nil

}

func (s service) GetProfileList(ctx context.Context) ([]models.Profile, error) {
	profiles, err := s.db.GetProfileList(ctx)
	if err != nil {
		return nil, fmt.Errorf("get profile list error: %w", err)
	}

	return profiles, nil

}

func (s service) GetAuthKey(ctx context.Context, authKey string) (bool, error) {
	ok, err := s.db.GetAuthKey(ctx, authKey)
	if err != nil {
		return false, fmt.Errorf("error get authKey: %w", err)
	}

	return ok, nil
}
