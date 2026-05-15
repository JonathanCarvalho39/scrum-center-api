package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type InviteRepository struct {
	invites map[uuid.UUID]*entity.Invite
	codes   map[string]*entity.Invite // For quick lookup by code
	mu      sync.RWMutex
}

func NewInviteRepository() *InviteRepository {
	return &InviteRepository{
		invites: make(map[uuid.UUID]*entity.Invite),
		codes:   make(map[string]*entity.Invite),
	}
}

func (r *InviteRepository) Create(ctx context.Context, invite *entity.Invite) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.invites[invite.ID] = invite
	r.codes[invite.Code] = invite
	return nil
}

func (r *InviteRepository) FindByCode(ctx context.Context, code string) (*entity.Invite, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	invite, exists := r.codes[code]
	if !exists {
		return nil, errors.New("invite not found")
	}

	return invite, nil
}

func (r *InviteRepository) FindByTeamID(ctx context.Context, teamID uuid.UUID) ([]*entity.Invite, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var invites []*entity.Invite
	for _, invite := range r.invites {
		if invite.TeamID == teamID {
			invites = append(invites, invite)
		}
	}

	return invites, nil
}
