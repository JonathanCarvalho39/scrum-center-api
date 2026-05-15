package invite

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/repository"
	"github.com/google/uuid"
)

type GenerateInviteInput struct {
	TeamID string
}

type GenerateInviteOutput struct {
	InviteID string
	TeamID   string
	Code     string
}

type GenerateInviteUseCase struct {
	teamRepo   repository.TeamRepository
	inviteRepo repository.InviteRepository
}

func NewGenerateInviteUseCase(teamRepo repository.TeamRepository, inviteRepo repository.InviteRepository) *GenerateInviteUseCase {
	return &GenerateInviteUseCase{
		teamRepo:   teamRepo,
		inviteRepo: inviteRepo,
	}
}

func (uc *GenerateInviteUseCase) Execute(ctx context.Context, input GenerateInviteInput) (*GenerateInviteOutput, error) {
	// Parse team ID
	teamID, err := uuid.Parse(input.TeamID)
	if err != nil {
		return nil, err
	}

	// Verify team exists
	_, err = uc.teamRepo.FindByID(ctx, teamID)
	if err != nil {
		return nil, err
	}

	// Create invite
	invite, err := entity.NewInvite(teamID)
	if err != nil {
		return nil, err
	}

	// Save invite
	if err := uc.inviteRepo.Create(ctx, invite); err != nil {
		return nil, err
	}

	return &GenerateInviteOutput{
		InviteID: invite.ID.String(),
		TeamID:   invite.TeamID.String(),
		Code:     invite.Code,
	}, nil
}
