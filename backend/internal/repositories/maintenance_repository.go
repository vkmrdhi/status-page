package repositories

import (
	"errors"
	"time"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"

	"gorm.io/gorm"
)

type MaintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

func (r *MaintenanceRepository) Create(maintenance *models.Maintenance) error {
	return r.db.Create(maintenance).Error
}

func (r *MaintenanceRepository) GetByID(id uint) (*models.Maintenance, error) {
	var maintenance models.Maintenance
	result := r.db.First(&maintenance, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("maintenance not found")
		}
		return nil, result.Error
	}
	return &maintenance, nil
}

func (r *MaintenanceRepository) ListByOrganization(orgID uint, status *models.MaintenanceStatus) ([]models.Maintenance, error) {
	var maintenances []models.Maintenance

	query := r.db.Where("organization_id = ?", orgID)

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	result := query.
		Order("scheduled_start DESC").
		Find(&maintenances)

	return maintenances, result.Error
}

func (r *MaintenanceRepository) Update(maintenance *models.Maintenance) error {
	return r.db.Save(maintenance).Error
}

func (r *MaintenanceRepository) GetActiveMaintenances() ([]models.Maintenance, error) {
	var maintenances []models.Maintenance

	now := time.Now()

	// Fetch ongoing or upcoming maintenances
	result := r.db.
		Where("status IN (?, ?) AND scheduled_start <= ? AND scheduled_end >= ?",
			models.MaintenanceScheduled,
			models.MaintenanceInProgress,
			now,
			now).
		Order("scheduled_start ASC").
		Find(&maintenances)

	return maintenances, result.Error
}

func (r *MaintenanceRepository) StartMaintenance(maintenanceID uint) error {
	now := time.Now()
	return r.db.Model(&models.Maintenance{}).
		Where("id = ?", maintenanceID).
		Updates(map[string]interface{}{
			"status":       models.MaintenanceInProgress,
			"actual_start": now,
		}).Error
}

func (r *MaintenanceRepository) CompleteMaintenance(maintenanceID uint) error {
	now := time.Now()
	return r.db.Model(&models.Maintenance{}).
		Where("id = ?", maintenanceID).
		Updates(map[string]interface{}{
			"status":     models.MaintenanceCompleted,
			"actual_end": now,
		}).Error
}

func (r *MaintenanceRepository) GetUpcomingMaintenances(orgID uint) ([]models.Maintenance, error) {
	now := time.Now()

	var maintenances []models.Maintenance

	result := r.db.
		Where("organization_id = ? AND status = ? AND scheduled_start > ?",
			orgID,
			models.MaintenanceScheduled,
			now).
		Order("scheduled_start ASC").
		Find(&maintenances)

	return maintenances, result.Error
}
