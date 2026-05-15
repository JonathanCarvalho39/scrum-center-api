package main

import (
	"log"
	"os"

	"github.com/JonathanCarvalho39/scrum-center-api/internal/handler"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/infrastructure/repository/memory"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/invite"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/member"
	"github.com/JonathanCarvalho39/scrum-center-api/internal/usecase/team"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Scrum Center API
// @version 1.0
// @description API REST para gerenciamento de projetos Scrum
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email jonathan.carvalho@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set Gin mode
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	// Initialize repositories (in-memory)
	teamRepo := memory.NewTeamRepository()
	roleRepo := memory.NewRoleRepository()
	inviteRepo := memory.NewInviteRepository()
	memberRepo := memory.NewMemberRepository()

	// Initialize use cases
	createTeamUseCase := team.NewCreateTeamUseCase(teamRepo, roleRepo)
	addCustomRoleUseCase := team.NewAddCustomRoleUseCase(teamRepo, roleRepo)
	generateInviteUseCase := invite.NewGenerateInviteUseCase(teamRepo, inviteRepo)
	joinTeamUseCase := member.NewJoinTeamUseCase(teamRepo, roleRepo, inviteRepo, memberRepo)

	// Initialize handlers
	teamHandler := handler.NewTeamHandler(createTeamUseCase, addCustomRoleUseCase)
	inviteHandler := handler.NewInviteHandler(generateInviteUseCase)
	memberHandler := handler.NewMemberHandler(joinTeamUseCase)

	// Initialize router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "scrum-center-api",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Team routes
		teams := v1.Group("/teams")
		{
			teams.POST("", teamHandler.CreateTeam)
			teams.POST("/roles", teamHandler.AddCustomRole)
		}

		// Invite routes
		invites := v1.Group("/invites")
		{
			invites.POST("", inviteHandler.GenerateInvite)
		}

		// Member routes
		members := v1.Group("/members")
		{
			members.POST("/join", memberHandler.JoinTeam)
		}

		// Legacy ping endpoint
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
