package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID           uuid.UUID
	TeamID       uuid.UUID
	Name         string
	IsPredefined bool
	CreatedAt    time.Time
}

func NewRole(teamID uuid.UUID, name string, isPredefined bool) (*Role, error) {
	if name == "" {
		return nil, ErrRoleNameRequired
	}

	return &Role{
		ID:           uuid.New(),
		TeamID:       teamID,
		Name:         name,
		IsPredefined: isPredefined,
		CreatedAt:    time.Now(),
	}, nil
}

// Predefined roles constants
const (
	RoleScrumMaster  = "Scrum Master"
	RoleProductOwner = "Product Owner"
	RoleDeveloper    = "Developer"
)

var PredefinedRoles = []string{
	RoleScrumMaster,
	RoleProductOwner,
	RoleDeveloper,
}
