package repository

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type TeamRepository interface {
	Create(ctx context.Context, team *entity.Team) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Team, error)
}
