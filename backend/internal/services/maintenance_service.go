package services

import (
	"errors"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/repositories"
)

// MaintenanceService handles maintenance-related business logic
type MaintenanceService struct {
	maintenanceRepo *repositories.MaintenanceRepository
	serviceRepo     *repositories.ServiceRepository
}

func NewMaintenanceService(maintenanceRepo *repositories.MaintenanceRepository, serviceRepo *repositories.ServiceRepository) *MaintenanceService {
	return &MaintenanceService{
		maintenanceRepo: maintenanceRepo,
		serviceRepo:     serviceRepo,
	}
}

func (s *MaintenanceService) ScheduleMaintenance(maintenance *models.Maintenance) error {
	// Validate all services exist
	for _, serviceID := range maintenance.ServiceIDs {
		_, err := s.serviceRepo.GetByID(serviceID)
		if err != nil {
			return errors.New("one or more services not found")
		}
	}

	maintenance.Status = models.MaintenanceScheduled
	return s.maintenanceRepo.Create(maintenance)
}

func (s *MaintenanceService) StartMaintenance(maintenanceID uint) error {
	return s.maintenanceRepo.StartMaintenance(maintenanceID)
}

func (s *MaintenanceService) CompleteMaintenance(maintenanceID uint) error {
	return s.maintenanceRepo.CompleteMaintenance(maintenanceID)
}
