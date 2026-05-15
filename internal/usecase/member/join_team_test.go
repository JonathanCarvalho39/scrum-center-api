package member_test

import (
	"context"
	"testing"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/domain/entity"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/infrastructure/repository/memory"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/member"
)

func TestJoinTeam_Success(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	inviteRepo := memory.NewInviteRepository()
	memberRepo := memory.NewMemberRepository()

	// Create team, role, and invite
	testTeam, _ := entity.NewTeam("Test Team")
	teamRepo.Create(context.Background(), testTeam)

	testRole, _ := entity.NewRole(testTeam.ID, "Developer", true)
	roleRepo.Create(context.Background(), testRole)

	testInvite, _ := entity.NewInvite(testTeam.ID)
	inviteRepo.Create(context.Background(), testInvite)

	useCase := member.NewJoinTeamUseCase(teamRepo, roleRepo, inviteRepo, memberRepo)

	input := member.JoinTeamInput{
		InviteCode: testInvite.Code,
		MemberName: "John Doe",
		RoleID:     testRole.ID.String(),
	}

	// Act
	output, err := useCase.Execute(context.Background(), input)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if output.MemberID == "" {
		t.Error("Expected MemberID to be set")
	}

	if output.MemberName != "John Doe" {
		t.Errorf("Expected member name 'John Doe', got: %s", output.MemberName)
	}

	if output.TeamID != testTeam.ID.String() {
		t.Errorf("Expected TeamID %s, got: %s", testTeam.ID.String(), output.TeamID)
	}

	if output.RoleName != "Developer" {
		t.Errorf("Expected role name 'Developer', got: %s", output.RoleName)
	}
}

func TestJoinTeam_InvalidInviteCode(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	inviteRepo := memory.NewInviteRepository()
	memberRepo := memory.NewMemberRepository()

	useCase := member.NewJoinTeamUseCase(teamRepo, roleRepo, inviteRepo, memberRepo)

	input := member.JoinTeamInput{
		InviteCode: "INVALID",
		MemberName: "John Doe",
		RoleID:     "some-role-id",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for invalid invite code, got nil")
	}
}

func TestJoinTeam_EmptyMemberName(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	inviteRepo := memory.NewInviteRepository()
	memberRepo := memory.NewMemberRepository()

	testTeam, _ := entity.NewTeam("Test Team")
	teamRepo.Create(context.Background(), testTeam)

	testInvite, _ := entity.NewInvite(testTeam.ID)
	inviteRepo.Create(context.Background(), testInvite)

	useCase := member.NewJoinTeamUseCase(teamRepo, roleRepo, inviteRepo, memberRepo)

	input := member.JoinTeamInput{
		InviteCode: testInvite.Code,
		MemberName: "",
		RoleID:     "some-role-id",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for empty member name, got nil")
	}

	if err != entity.ErrMemberNameRequired {
		t.Errorf("Expected ErrMemberNameRequired, got: %v", err)
	}
}

func TestJoinTeam_DuplicateMemberName(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	inviteRepo := memory.NewInviteRepository()
	memberRepo := memory.NewMemberRepository()

	testTeam, _ := entity.NewTeam("Test Team")
	teamRepo.Create(context.Background(), testTeam)

	testRole, _ := entity.NewRole(testTeam.ID, "Developer", true)
	roleRepo.Create(context.Background(), testRole)

	testInvite, _ := entity.NewInvite(testTeam.ID)
	inviteRepo.Create(context.Background(), testInvite)

	// Create existing member
	existingMember, _ := entity.NewMember(testTeam.ID, testRole.ID, "John Doe")
	memberRepo.Create(context.Background(), existingMember)

	useCase := member.NewJoinTeamUseCase(teamRepo, roleRepo, inviteRepo, memberRepo)

	input := member.JoinTeamInput{
		InviteCode: testInvite.Code,
		MemberName: "John Doe",
		RoleID:     testRole.ID.String(),
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for duplicate member name, got nil")
	}

	if err != entity.ErrMemberAlreadyExists {
		t.Errorf("Expected ErrMemberAlreadyExists, got: %v", err)
	}
}

func TestJoinTeam_InvalidRoleID(t *testing.T) {
	// Arrange
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	inviteRepo := memory.NewInviteRepository()
	memberRepo := memory.NewMemberRepository()

	testTeam, _ := entity.NewTeam("Test Team")
	teamRepo.Create(context.Background(), testTeam)

	testInvite, _ := entity.NewInvite(testTeam.ID)
	inviteRepo.Create(context.Background(), testInvite)

	useCase := member.NewJoinTeamUseCase(teamRepo, roleRepo, inviteRepo, memberRepo)

	input := member.JoinTeamInput{
		InviteCode: testInvite.Code,
		MemberName: "John Doe",
		RoleID:     "invalid-uuid",
	}

	// Act
	_, err := useCase.Execute(context.Background(), input)

	// Assert
	if err == nil {
		t.Fatal("Expected error for invalid role ID, got nil")
	}
}
