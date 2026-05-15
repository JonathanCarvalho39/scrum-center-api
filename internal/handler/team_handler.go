package handler

import (
	"net/http"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/team"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	createTeamUseCase     *team.CreateTeamUseCase
	addCustomRoleUseCase  *team.AddCustomRoleUseCase
}

func NewTeamHandler(
	createTeamUseCase *team.CreateTeamUseCase,
	addCustomRoleUseCase *team.AddCustomRoleUseCase,
) *TeamHandler {
	return &TeamHandler{
		createTeamUseCase:    createTeamUseCase,
		addCustomRoleUseCase: addCustomRoleUseCase,
	}
}

type CreateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateTeamResponse struct {
	TeamID          string              `json:"team_id"`
	TeamName        string              `json:"team_name"`
	PredefinedRoles []RoleResponse      `json:"predefined_roles"`
}

type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateTeam godoc
// @Summary Create a new team
// @Description Creates a new team with predefined roles (Scrum Master, Product Owner, Developer)
// @Tags teams
// @Accept json
// @Produce json
// @Param request body CreateTeamRequest true "Team creation request"
// @Success 201 {object} CreateTeamResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teams [post]
func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var req CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	input := team.CreateTeamInput{
		Name: req.Name,
	}

	output, err := h.createTeamUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	var roles []RoleResponse
	for _, role := range output.PredefinedRoles {
		roles = append(roles, RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		})
	}

	c.JSON(http.StatusCreated, CreateTeamResponse{
		TeamID:          output.TeamID,
		TeamName:        output.TeamName,
		PredefinedRoles: roles,
	})
}

type AddCustomRoleRequest struct {
	TeamID   string `json:"team_id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

type AddCustomRoleResponse struct {
	RoleID       string `json:"role_id"`
	RoleName     string `json:"role_name"`
	IsPredefined bool   `json:"is_predefined"`
}

// AddCustomRole godoc
// @Summary Add a custom role to a team
// @Description Allows adding a custom role to an existing team
// @Tags teams
// @Accept json
// @Produce json
// @Param request body AddCustomRoleRequest true "Custom role request"
// @Success 201 {object} AddCustomRoleResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /teams/roles [post]
func (h *TeamHandler) AddCustomRole(c *gin.Context) {
	var req AddCustomRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	input := team.AddCustomRoleInput{
		TeamID:   req.TeamID,
		RoleName: req.RoleName,
	}

	output, err := h.addCustomRoleUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		// TODO: Differentiate between 404 and 500 errors
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, AddCustomRoleResponse{
		RoleID:       output.RoleID,
		RoleName:     output.RoleName,
		IsPredefined: output.IsPredefined,
	})
}

