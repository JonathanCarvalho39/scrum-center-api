package repository

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type RoleRepository interface {
	Create(ctx context.Context, role *entity.Role) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	FindByTeamID(ctx context.Context, teamID uuid.UUID) ([]*entity.Role, error)
}
