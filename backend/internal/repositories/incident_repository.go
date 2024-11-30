package repositories

import (
	"errors"
	"time"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"

	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{db: db}
}

func (r *IncidentRepository) Create(incident *models.Incident) error {
	return r.db.Create(incident).Error
}

func (r *IncidentRepository) GetByID(id uint) (*models.Incident, error) {
	var incident models.Incident
	result := r.db.
		Preload("Service").
		Preload("Updates").
		First(&incident, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("incident not found")
		}
		return nil, result.Error
	}
	return &incident, nil
}

func (r *IncidentRepository) ListByOrganization(orgID uint, status *models.IncidentStatus) ([]models.Incident, error) {
	var incidents []models.Incident

	query := r.db.Where("organization_id = ?", orgID)

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	result := query.
		Preload("Service").
		Preload("Updates").
		Order("created_at DESC").
		Find(&incidents)

	return incidents, result.Error
}

func (r *IncidentRepository) Update(incident *models.Incident) error {
	return r.db.Save(incident).Error
}

func (r *IncidentRepository) AddUpdate(update *models.IncidentUpdate) error {
	return r.db.Create(update).Error
}

func (r *IncidentRepository) GetActiveIncidents() ([]models.Incident, error) {
	var incidents []models.Incident

	// Fetch incidents that are not resolved
	result := r.db.
		Where("status != ?", models.IncidentResolved).
		Preload("Service").
		Preload("Updates", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Order("created_at DESC").
		Find(&incidents)

	return incidents, result.Error
}

func (r *IncidentRepository) ResolveIncident(incidentID uint) error {
	return r.db.Model(&models.Incident{}).
		Where("id = ?", incidentID).
		Updates(map[string]interface{}{
			"status":      models.IncidentResolved,
			"resolved_at": time.Now(),
		}).Error
}

func (r *IncidentRepository) GetRecentIncidents(orgID uint, days int) ([]models.Incident, error) {
	var incidents []models.Incident

	result := r.db.
		Where("organization_id = ? AND created_at >= ?", orgID, time.Now().AddDate(0, 0, -days)).
		Preload("Service").
		Preload("Updates", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at DESC")
		}).
		Order("created_at DESC").
		Find(&incidents)

	return incidents, result.Error
}
