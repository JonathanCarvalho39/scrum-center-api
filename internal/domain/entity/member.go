package entity

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID       uuid.UUID
	TeamID   uuid.UUID
	RoleID   uuid.UUID
	Name     string
	JoinedAt time.Time
}

func NewMember(teamID, roleID uuid.UUID, name string) (*Member, error) {
	if name == "" {
		return nil, ErrMemberNameRequired
	}

	return &Member{
		ID:       uuid.New(),
		TeamID:   teamID,
		RoleID:   roleID,
		Name:     name,
		JoinedAt: time.Now(),
	}, nil
}
