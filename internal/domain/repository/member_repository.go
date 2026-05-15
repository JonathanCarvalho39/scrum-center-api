package repository

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type MemberRepository interface {
	Create(ctx context.Context, member *entity.Member) error
	FindByTeamID(ctx context.Context, teamID uuid.UUID) ([]*entity.Member, error)
	FindByTeamAndName(ctx context.Context, teamID uuid.UUID, name string) (*entity.Member, error)
}

