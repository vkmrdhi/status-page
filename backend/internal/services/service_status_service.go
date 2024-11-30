package services

import (
	"time"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/repositories"
)

type ServiceStatusService struct {
	serviceRepo      *repositories.ServiceRepository
	incidentRepo     *repositories.IncidentRepository
	organizationRepo *repositories.OrganizationRepository
}

func NewServiceStatusService(serviceRepo *repositories.ServiceRepository, incidentRepo *repositories.IncidentRepository) *ServiceStatusService {
	return &ServiceStatusService{
		serviceRepo:  serviceRepo,
		incidentRepo: incidentRepo,
	}
}

func (s *ServiceStatusService) UpdateServiceStatus(serviceID uint, newStatus models.ServiceStatus, incidentDetails *models.Incident) error {
	// Update service status
	if err := s.serviceRepo.UpdateStatus(serviceID, newStatus); err != nil {
		return err
	}

	// If status indicates an issue, create an incident
	if newStatus != models.StatusOperational && incidentDetails != nil {
		incidentDetails.ServiceID = serviceID
		incidentDetails.Status = models.IncidentInvestigating
		incidentDetails.StartedAt = time.Now()

		if err := s.incidentRepo.Create(incidentDetails); err != nil {
			return err
		}
	}

	return nil
}

func (s *ServiceStatusService) GetServiceStatusSummary(orgID uint) (map[models.ServiceStatus]int, error) {
	services, err := s.serviceRepo.ListByOrganization(orgID)
	if err != nil {
		return nil, err
	}

	statusSummary := make(map[models.ServiceStatus]int)
	for _, service := range services {
		statusSummary[service.Status]++
	}

	return statusSummary, nil
}
