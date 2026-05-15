package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type TeamRepository struct {
	teams      map[uuid.UUID]*entity.Team
	mu         sync.RWMutex
	forceError bool
}

func NewTeamRepository() *TeamRepository {
	return &TeamRepository{
		teams: make(map[uuid.UUID]*entity.Team),
	}
}

func NewTeamRepositoryWithError() *TeamRepository {
	return &TeamRepository{
		teams:      make(map[uuid.UUID]*entity.Team),
		forceError: true,
	}
}

func (r *TeamRepository) Create(ctx context.Context, team *entity.Team) error {
	if r.forceError {
		return errors.New("repository error")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.teams[team.ID] = team
	return nil
}

func (r *TeamRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Team, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	team, exists := r.teams[id]
	if !exists {
		return nil, errors.New("team not found")
	}

	return team, nil
}
