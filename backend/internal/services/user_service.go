package services

import (
	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
	orgRepo  *repositories.OrganizationRepository
}

func NewUserService(userRepo *repositories.UserRepository, orgRepo *repositories.OrganizationRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		orgRepo:  orgRepo,
	}
}

func (s *UserService) CreateUserWithOrganization(email, name, auth0ID, orgName, subdomain string) (*models.User, error) {
	// Create organization first
	org := &models.Organization{
		Name:      orgName,
		Subdomain: subdomain,
	}
	if err := s.orgRepo.Create(org); err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Email:          email,
		Name:           name,
		Auth0ID:        auth0ID,
		OrganizationID: org.ID,
		Role:           models.RoleAdmin,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
