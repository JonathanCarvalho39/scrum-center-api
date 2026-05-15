package repository

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type InviteRepository interface {
	Create(ctx context.Context, invite *entity.Invite) error
	FindByCode(ctx context.Context, code string) (*entity.Invite, error)
	FindByTeamID(ctx context.Context, teamID uuid.UUID) ([]*entity.Invite, error)
}

