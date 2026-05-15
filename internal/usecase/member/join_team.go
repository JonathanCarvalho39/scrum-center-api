package member

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/repository"
	"github.com/google/uuid"
)

type JoinTeamInput struct {
	InviteCode string
	MemberName string
	RoleID     string
}

type JoinTeamOutput struct {
	MemberID   string
	MemberName string
	TeamID     string
	TeamName   string
	RoleID     string
	RoleName   string
}

type JoinTeamUseCase struct {
	teamRepo   repository.TeamRepository
	roleRepo   repository.RoleRepository
	inviteRepo repository.InviteRepository
	memberRepo repository.MemberRepository
}

func NewJoinTeamUseCase(
	teamRepo repository.TeamRepository,
	roleRepo repository.RoleRepository,
	inviteRepo repository.InviteRepository,
	memberRepo repository.MemberRepository,
) *JoinTeamUseCase {
	return &JoinTeamUseCase{
		teamRepo:   teamRepo,
		roleRepo:   roleRepo,
		inviteRepo: inviteRepo,
		memberRepo: memberRepo,
	}
}

func (uc *JoinTeamUseCase) Execute(ctx context.Context, input JoinTeamInput) (*JoinTeamOutput, error) {
	// Validate member name early
	if input.MemberName == "" {
		return nil, entity.ErrMemberNameRequired
	}

	// Find invite by code
	invite, err := uc.inviteRepo.FindByCode(ctx, input.InviteCode)
	if err != nil {
		return nil, err
	}

	// Check if member already exists in this team
	existingMember, _ := uc.memberRepo.FindByTeamAndName(ctx, invite.TeamID, input.MemberName)
	if existingMember != nil {
		return nil, entity.ErrMemberAlreadyExists
	}

	// Parse role ID
	roleID, err := uuid.Parse(input.RoleID)
	if err != nil {
		return nil, err
	}

	// Get role details
	role, err := uc.roleRepo.FindByID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	// Get team details
	team, err := uc.teamRepo.FindByID(ctx, invite.TeamID)
	if err != nil {
		return nil, err
	}

	// Create member
	newMember, err := entity.NewMember(invite.TeamID, roleID, input.MemberName)
	if err != nil {
		return nil, err
	}

	// Save member
	if err := uc.memberRepo.Create(ctx, newMember); err != nil {
		return nil, err
	}

	return &JoinTeamOutput{
		MemberID:   newMember.ID.String(),
		MemberName: newMember.Name,
		TeamID:     team.ID.String(),
		TeamName:   team.Name,
		RoleID:     role.ID.String(),
		RoleName:   role.Name,
	}, nil
}
