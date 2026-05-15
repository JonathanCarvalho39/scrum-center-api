package team

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/repository"
	"github.com/google/uuid"
)

type AddCustomRoleInput struct {
	TeamID   string
	RoleName string
}

type AddCustomRoleOutput struct {
	RoleID       string
	RoleName     string
	IsPredefined bool
}

type AddCustomRoleUseCase struct {
	teamRepo repository.TeamRepository
	roleRepo repository.RoleRepository
}

func NewAddCustomRoleUseCase(teamRepo repository.TeamRepository, roleRepo repository.RoleRepository) *AddCustomRoleUseCase {
	return &AddCustomRoleUseCase{
		teamRepo: teamRepo,
		roleRepo: roleRepo,
	}
}

func (uc *AddCustomRoleUseCase) Execute(ctx context.Context, input AddCustomRoleInput) (*AddCustomRoleOutput, error) {
	// Validate role name first (fail fast)
	if input.RoleName == "" {
		return nil, entity.ErrRoleNameRequired
	}

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

	// Create custom role
	role, err := entity.NewRole(teamID, input.RoleName, false)
	if err != nil {
		return nil, err
	}

	// Save role
	if err := uc.roleRepo.Create(ctx, role); err != nil {
		return nil, err
	}

	return &AddCustomRoleOutput{
		RoleID:       role.ID.String(),
		RoleName:     role.Name,
		IsPredefined: role.IsPredefined,
	}, nil
}


