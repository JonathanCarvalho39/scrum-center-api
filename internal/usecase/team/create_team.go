package team

import (
	"context"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/repository"
)

type CreateTeamInput struct {
	Name string
}

type RoleOutput struct {
	ID   string
	Name string
}

type CreateTeamOutput struct {
	TeamID          string
	TeamName        string
	PredefinedRoles []RoleOutput
}

type CreateTeamUseCase struct {
	teamRepo repository.TeamRepository
	roleRepo repository.RoleRepository
}

func NewCreateTeamUseCase(teamRepo repository.TeamRepository, roleRepo repository.RoleRepository) *CreateTeamUseCase {
	return &CreateTeamUseCase{
		teamRepo: teamRepo,
		roleRepo: roleRepo,
	}
}

func (uc *CreateTeamUseCase) Execute(ctx context.Context, input CreateTeamInput) (*CreateTeamOutput, error) {
	// Create team
	team, err := entity.NewTeam(input.Name)
	if err != nil {
		return nil, err
	}

	// Save team
	if err := uc.teamRepo.Create(ctx, team); err != nil {
		return nil, err
	}

	// Create predefined roles
	var rolesOutput []RoleOutput
	for _, roleName := range entity.PredefinedRoles {
		role, err := entity.NewRole(team.ID, roleName, true)
		if err != nil {
			return nil, err
		}

		if err := uc.roleRepo.Create(ctx, role); err != nil {
			return nil, err
		}

		rolesOutput = append(rolesOutput, RoleOutput{
			ID:   role.ID.String(),
			Name: role.Name,
		})
	}

	return &CreateTeamOutput{
		TeamID:          team.ID.String(),
		TeamName:        team.Name,
		PredefinedRoles: rolesOutput,
	}, nil
}
