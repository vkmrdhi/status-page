package repositories

import (
	"errors"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) Create(org *models.Organization) error {
	return r.db.Create(org).Error
}

func (r *OrganizationRepository) GetByID(id uint) (*models.Organization, error) {
	var org models.Organization
	result := r.db.First(&org, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, result.Error
	}
	return &org, nil
}

func (r *OrganizationRepository) GetBySlug(slug string) (*models.Organization, error) {
	var org models.Organization
	result := r.db.Where("slug = ?", slug).First(&org)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("organization not found")
		}
		return nil, result.Error
	}
	return &org, nil
}

func (r *OrganizationRepository) ListUserOrganizations(userID string) ([]models.Organization, error) {
	var organizations []models.Organization

	// This is a complex query to fetch organizations where the user is a member
	result := r.db.
		Joins("JOIN user_organizations ON user_organizations.organization_id = organizations.id").
		Joins("JOIN users ON users.id = user_organizations.user_id").
		Where("users.auth0_id = ?", userID).
		Find(&organizations)

	return organizations, result.Error
}

func (r *OrganizationRepository) GetUserOrganization(userID string, orgID uint) (*models.Organization, error) {
	var org models.Organization

	result := r.db.
		Joins("JOIN user_organizations ON user_organizations.organization_id = organizations.id").
		Joins("JOIN users ON users.id = user_organizations.user_id").
		Where("users.auth0_id = ? AND organizations.id = ?", userID, orgID).
		First(&org)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user does not have access to this organization")
		}
		return nil, result.Error
	}

	return &org, nil
}

func (r *OrganizationRepository) Update(org *models.Organization) error {
	return r.db.Save(org).Error
}

func (r *OrganizationRepository) Delete(id uint) error {
	return r.db.Delete(&models.Organization{}, id).Error
}
