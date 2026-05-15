package team_test

import (
	"context"
	"testing"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/infrastructure/repository/memory"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/team"
)

func TestCreateTeam_Success(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	useCase := team.NewCreateTeamUseCase(teamRepo, roleRepo)

	input := team.CreateTeamInput{
		Name: "My Scrum Team",
	}

	// Act
	output, err := useCase.Execute(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if output.TeamID == "" {
		t.Error("Expected TeamID to be set")
	}

	if output.TeamName != "My Scrum Team" {
		t.Errorf("Expected team name 'My Scrum Team', got: %s", output.TeamName)
	}

	if len(output.PredefinedRoles) != 3 {
		t.Errorf("Expected 3 predefined roles, got: %d", len(output.PredefinedRoles))
	}

	// Verify predefined roles were created
	expectedRoles := map[string]bool{
		entity.RoleScrumMaster:  false,
		entity.RoleProductOwner: false,
		entity.RoleDeveloper:    false,
	}

	for _, role := range output.PredefinedRoles {
		if _, exists := expectedRoles[role.Name]; exists {
			expectedRoles[role.Name] = true
		}
	}

	for roleName, found := range expectedRoles {
		if !found {
			t.Errorf("Expected predefined role '%s' not found", roleName)
		}
	}
}

func TestCreateTeam_EmptyName(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	useCase := team.NewCreateTeamUseCase(teamRepo, roleRepo)

	input := team.CreateTeamInput{
		Name: "",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for empty team name, got nil")
	}

	if err != entity.ErrTeamNameRequired {
		t.Errorf("Expected ErrTeamNameRequired, got: %v", err)
	}
}

func TestCreateTeam_RepositoryError(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepositoryWithError()
	roleRepo := memory.NewRoleRepository()
	useCase := team.NewCreateTeamUseCase(teamRepo, roleRepo)

	input := team.CreateTeamInput{
		Name: "Test Team",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error from repository, got nil")
	}
}
