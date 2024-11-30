package services

import (
	"errors"
	"time"

	"github.com/vikasatfactors/status-page-app/backend/internal/models"
	"github.com/vikasatfactors/status-page-app/backend/internal/repositories"
)

type IncidentService struct {
	incidentRepo *repositories.IncidentRepository
	serviceRepo  *repositories.ServiceRepository
}

func NewIncidentService(incidentRepo *repositories.IncidentRepository, serviceRepo *repositories.ServiceRepository) *IncidentService {
	return &IncidentService{
		incidentRepo: incidentRepo,
		serviceRepo:  serviceRepo,
	}
}

func (s *IncidentService) CreateIncident(incident *models.Incident) error {
	// Validate service exists
	_, err := s.serviceRepo.GetServiceByID(incident.ServiceID)
	if err != nil {
		return errors.New("service not found")
	}

	// Set initial status and start time
	incident.Status = models.IncidentInvestigating
	incident.StartedAt = time.Now()

	return s.incidentRepo.Create(incident)
}

func (s *IncidentService) AddIncidentUpdate(incidentID uint, update *models.IncidentUpdate) error {
	// Update incident status based on the update
	return s.incidentRepo.AddUpdate(incidentID, update)
}

func (s *IncidentService) ResolveIncident(incidentID uint) error {
	return s.incidentRepo.ResolveIncident(incidentID)
}
