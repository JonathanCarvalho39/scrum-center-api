package invite_test

import (
	"context"
	"testing"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/infrastructure/repository/memory"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/invite"
	"github.com/google/uuid"
)

func TestGenerateInvite_Success(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	inviteRepo := memory.NewInviteRepository()

	// Create a team first
	testTeam, _ := entity.NewTeam("Test Team")
	teamRepo.Create(context.Background(), testTeam)

	useCase := invite.NewGenerateInviteUseCase(teamRepo, inviteRepo)

	input := invite.GenerateInviteInput{
		TeamID: testTeam.ID.String(),
	}

	// Act
	output, err := useCase.Execute(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if output.InviteID == "" {
		t.Error("Expected InviteID to be set")
	}

	if output.Code == "" {
		t.Error("Expected invite code to be set")
	}

	if len(output.Code) < 5 {
		t.Errorf("Expected invite code to have at least 5 characters, got: %s", output.Code)
	}

	if output.TeamID != testTeam.ID.String() {
		t.Errorf("Expected TeamID %s, got: %s", testTeam.ID.String(), output.TeamID)
	}
}

func TestGenerateInvite_TeamNotFound(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	inviteRepo := memory.NewInviteRepository()
	useCase := invite.NewGenerateInviteUseCase(teamRepo, inviteRepo)

	input := invite.GenerateInviteInput{
		TeamID: uuid.New().String(),
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for non-existent team, got nil")
	}
}

func TestGenerateInvite_InvalidTeamID(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	inviteRepo := memory.NewInviteRepository()
	useCase := invite.NewGenerateInviteUseCase(teamRepo, inviteRepo)

	input := invite.GenerateInviteInput{
		TeamID: "invalid-uuid",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for invalid UUID, got nil")
	}
}

func TestGenerateInvite_MultipleInvites(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	inviteRepo := memory.NewInviteRepository()

	// Create a team first
	testTeam, _ := entity.NewTeam("Test Team")
	teamRepo.Create(context.Background(), testTeam)

	useCase := invite.NewGenerateInviteUseCase(teamRepo, inviteRepo)

	input := invite.GenerateInviteInput{
		TeamID: testTeam.ID.String(),
	}

	// Act - Generate multiple invites
	output1, _ := useCase.Execute(context.Background(), input)
	output2, _ := useCase.Execute(context.Background(), input)

	// Assert - Codes should be unique
	if output1.Code == output2.Code {
		t.Error("Expected different invite codes, got the same code")
	}

	if output1.InviteID == output2.InviteID {
		t.Error("Expected different invite IDs, got the same ID")
	}
}

