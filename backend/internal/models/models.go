package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;not null"`
	Name           string
	Auth0ID        string `gorm:"uniqueIndex;not null"`
	OrganizationID uint
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	Role           UserRole     `gorm:"not null"`
}

// UserRole defines the different roles a user can have
type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleMember UserRole = "member"
	RoleViewer UserRole = "viewer"
)

// Organization represents a multi-tenant organization
type Organization struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Subdomain    string `gorm:"uniqueIndex;not null"`
	Users        []User
	Services     []Service
	Incidents    []Incident
	Maintenances []Maintenance
}

// ServiceStatus defines the possible statuses for a service
type ServiceStatus string

const (
	StatusOperational   ServiceStatus = "operational"
	StatusDegraded      ServiceStatus = "degraded"
	StatusPartialOutage ServiceStatus = "partial_outage"
	StatusMajorOutage   ServiceStatus = "major_outage"
)

// Service represents a monitored service within an organization
type Service struct {
	gorm.Model
	Name             string `gorm:"not null"`
	Description      string
	Status           ServiceStatus `gorm:"not null"`
	OrganizationID   uint
	Organization     Organization `gorm:"foreignKey:OrganizationID"`
	Incidents        []Incident   `gorm:"foreignKey:ServiceID"`
	LastStatusChange time.Time
}

// IncidentStatus defines the possible statuses for an incident
type IncidentStatus string

const (
	IncidentInvestigating IncidentStatus = "investigating"
	IncidentIdentified    IncidentStatus = "identified"
	IncidentMonitoring    IncidentStatus = "monitoring"
	IncidentResolved      IncidentStatus = "resolved"
)

// Incident represents a service disruption or issue
type Incident struct {
	gorm.Model
	Title          string `gorm:"not null"`
	Description    string
	Status         IncidentStatus `gorm:"not null"`
	ServiceID      uint
	Service        Service `gorm:"foreignKey:ServiceID"`
	OrganizationID uint
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	StartedAt      time.Time
	ResolvedAt     *time.Time
	Updates        []IncidentUpdate
}

// IncidentUpdate represents updates to an ongoing incident
type IncidentUpdate struct {
	gorm.Model
	IncidentID  uint     `gorm:"not null"`
	Incident    Incident `gorm:"foreignKey:IncidentID"`
	Description string   `gorm:"not null"`
	Status      IncidentStatus
	CreatedAt   time.Time
}

// Maintenance represents scheduled maintenance events
type Maintenance struct {
	gorm.Model
	Title          string `gorm:"not null"`
	Description    string
	Status         MaintenanceStatus `gorm:"not null"`
	ServiceIDs     []uint            `gorm:"-"`
	Services       []Service         `gorm:"many2many:maintenance_services;"`
	OrganizationID uint
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	ScheduledStart time.Time
	ScheduledEnd   time.Time
	ActualStart    *time.Time
	ActualEnd      *time.Time
}

// MaintenanceStatus defines the possible statuses for maintenance
type MaintenanceStatus string

const (
	MaintenanceScheduled  MaintenanceStatus = "scheduled"
	MaintenanceInProgress MaintenanceStatus = "in_progress"
	MaintenanceCompleted  MaintenanceStatus = "completed"
)
