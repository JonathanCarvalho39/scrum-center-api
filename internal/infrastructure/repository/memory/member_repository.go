package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/google/uuid"
)

type MemberRepository struct {
	members map[uuid.UUID]*entity.Member
	mu      sync.RWMutex
}

func NewMemberRepository() *MemberRepository {
	return &MemberRepository{
		members: make(map[uuid.UUID]*entity.Member),
	}
}

func (r *MemberRepository) Create(ctx context.Context, member *entity.Member) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.members[member.ID] = member
	return nil
}

func (r *MemberRepository) FindByTeamID(ctx context.Context, teamID uuid.UUID) ([]*entity.Member, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var members []*entity.Member
	for _, member := range r.members {
		if member.TeamID == teamID {
			members = append(members, member)
		}
	}

	return members, nil
}

func (r *MemberRepository) FindByTeamAndName(ctx context.Context, teamID uuid.UUID, name string) (*entity.Member, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, member := range r.members {
		if member.TeamID == teamID && member.Name == name {
			return member, nil
		}
	}

	return nil, errors.New("member not found")
}
