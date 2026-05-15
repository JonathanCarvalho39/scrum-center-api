package team_test

import (
	"context"
	"testing"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/infrastructure/repository/memory"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/team"
	"github.com/google/uuid"
)

func TestAddCustomRole_Success(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()

	// Create a team first
	testTeam, _ := entity.NewTeam("Test Team")
	teamRepo.Create(context.Background(), testTeam)

	useCase := team.NewAddCustomRoleUseCase(teamRepo, roleRepo)

	input := team.AddCustomRoleInput{
		TeamID:   testTeam.ID.String(),
		RoleName: "QA Tester",
	}

	// Act
	output, err := useCase.Execute(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if output.RoleID == "" {
		t.Error("Expected RoleID to be set")
	}

	if output.RoleName != "QA Tester" {
		t.Errorf("Expected role name 'QA Tester', got: %s", output.RoleName)
	}

	if output.IsPredefined {
		t.Error("Expected custom role to not be predefined")
	}
}

func TestAddCustomRole_EmptyRoleName(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	useCase := team.NewAddCustomRoleUseCase(teamRepo, roleRepo)

	input := team.AddCustomRoleInput{
		TeamID:   uuid.New().String(),
		RoleName: "",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for empty role name, got nil")
	}

	if err != entity.ErrRoleNameRequired {
		t.Errorf("Expected ErrRoleNameRequired, got: %v", err)
	}
}

func TestAddCustomRole_TeamNotFound(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	useCase := team.NewAddCustomRoleUseCase(teamRepo, roleRepo)

	input := team.AddCustomRoleInput{
		TeamID:   uuid.New().String(),
		RoleName: "Designer",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for non-existent team, got nil")
	}
}

func TestAddCustomRole_InvalidTeamID(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	useCase := team.NewAddCustomRoleUseCase(teamRepo, roleRepo)

	input := team.AddCustomRoleInput{
		TeamID:   "invalid-uuid",
		RoleName: "Designer",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for invalid UUID, got nil")
	}
}

