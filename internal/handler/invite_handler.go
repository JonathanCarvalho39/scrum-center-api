package handler

import (
	"net/http"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/invite"
	"github.com/gin-gonic/gin"
)

type InviteHandler struct {
	generateInviteUseCase *invite.GenerateInviteUseCase
}

func NewInviteHandler(generateInviteUseCase *invite.GenerateInviteUseCase) *InviteHandler {
	return &InviteHandler{
		generateInviteUseCase: generateInviteUseCase,
	}
}

type GenerateInviteRequest struct {
	TeamID string `json:"team_id" binding:"required"`
}

type GenerateInviteResponse struct {
	InviteID string `json:"invite_id"`
	TeamID   string `json:"team_id"`
	Code     string `json:"code"`
}

// GenerateInvite godoc
// @Summary Generate an invite code for a team
// @Description Creates a unique invite code that can be used to join the team
// @Tags invites
// @Accept json
// @Produce json
// @Param request body GenerateInviteRequest true "Generate invite request"
// @Success 201 {object} GenerateInviteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /invites [post]
func (h *InviteHandler) GenerateInvite(c *gin.Context) {
	var req GenerateInviteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	input := invite.GenerateInviteInput{
		TeamID: req.TeamID,
	}

	output, err := h.generateInviteUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, GenerateInviteResponse{
		InviteID: output.InviteID,
		TeamID:   output.TeamID,
		Code:     output.Code,
	})
}

