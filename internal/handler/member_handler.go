package handler

import (
	"net/http"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/member"
	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	joinTeamUseCase *member.JoinTeamUseCase
}

func NewMemberHandler(joinTeamUseCase *member.JoinTeamUseCase) *MemberHandler {
	return &MemberHandler{
		joinTeamUseCase: joinTeamUseCase,
	}
}

type JoinTeamRequest struct {
	InviteCode string `json:"invite_code" binding:"required"`
	MemberName string `json:"member_name" binding:"required"`
	RoleID     string `json:"role_id" binding:"required"`
}

type JoinTeamResponse struct {
	MemberID   string `json:"member_id"`
	MemberName string `json:"member_name"`
	TeamID     string `json:"team_id"`
	TeamName   string `json:"team_name"`
	RoleID     string `json:"role_id"`
	RoleName   string `json:"role_name"`
}

// JoinTeam godoc
// @Summary Join a team using an invite code
// @Description Allows a member to join a team by providing an invite code, name, and role
// @Tags members
// @Accept json
// @Produce json
// @Param request body JoinTeamRequest true "Join team request"
// @Success 201 {object} JoinTeamResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /members/join [post]
func (h *MemberHandler) JoinTeam(c *gin.Context) {
	var req JoinTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	input := member.JoinTeamInput{
		InviteCode: req.InviteCode,
		MemberName: req.MemberName,
		RoleID:     req.RoleID,
	}

	output, err := h.joinTeamUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		// TODO: Differentiate between 404, 409, and 500 errors
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, JoinTeamResponse{
		MemberID:   output.MemberID,
		MemberName: output.MemberName,
		TeamID:     output.TeamID,
		TeamName:   output.TeamName,
		RoleID:     output.RoleID,
		RoleName:   output.RoleName,
	})
}

