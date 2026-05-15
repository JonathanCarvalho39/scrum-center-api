package entity

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

func NewTeam(name string) (*Team, error) {
	if name == "" {
		return nil, ErrTeamNameRequired
	}

	return &Team{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}
