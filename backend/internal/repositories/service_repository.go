package repositories

import (
	"errors"
	"time"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"

	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) Create(service *models.Service) error {
	return r.db.Create(service).Error
}

func (r *ServiceRepository) GetByID(id uint) (*models.Service, error) {
	var service models.Service
	result := r.db.First(&service, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("service not found")
		}
		return nil, result.Error
	}
	return &service, nil
}

func (r *ServiceRepository) ListByOrganization(orgID uint) ([]models.Service, error) {
	var services []models.Service
	result := r.db.Where("organization_id = ?", orgID).Find(&services)
	return services, result.Error
}

func (r *ServiceRepository) Update(service *models.Service) error {
	// Update status and last status change
	service.LastStatusChange = time.Now()
	return r.db.Save(service).Error
}

func (r *ServiceRepository) UpdateStatus(serviceID uint, newStatus models.ServiceStatus) error {
	return r.db.Model(&models.Service{}).Where("id = ?", serviceID).Updates(map[string]interface{}{
		"status":             newStatus,
		"last_status_change": time.Now(),
	}).Error
}

func (r *ServiceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Service{}, id).Error
}

// GetPublicStatus retrieves services with their current status for public display
func (r *ServiceRepository) GetPublicStatus() ([]models.Service, error) {
	var services []models.Service
	result := r.db.Select("id", "name", "status", "last_status_change").Find(&services)
	return services, result.Error
}
