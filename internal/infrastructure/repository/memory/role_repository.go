package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type RoleRepository struct {
	roles map[uuid.UUID]*entity.Role
	mu    sync.RWMutex
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{
		roles: make(map[uuid.UUID]*entity.Role),
	}
}

func (r *RoleRepository) Create(ctx context.Context, role *entity.Role) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.roles[role.ID] = role
	return nil
}

func (r *RoleRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	role, exists := r.roles[id]
	if !exists {
		return nil, errors.New("role not found")
	}

	return role, nil
}

func (r *RoleRepository) FindByTeamID(ctx context.Context, teamID uuid.UUID) ([]*entity.Role, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var roles []*entity.Role
	for _, role := range r.roles {
		if role.TeamID == teamID {
			roles = append(roles, role)
		}
	}

	return roles, nil
}
