package entity

import (
	"crypto/rand"
	"encoding/base32"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Invite struct {
	ID        uuid.UUID
	TeamID    uuid.UUID
	Code      string
	CreatedAt time.Time
}

func NewInvite(teamID uuid.UUID) (*Invite, error) {
	code, err := generateInviteCode()
	if err != nil {
		return nil, err
	}

	return &Invite{
		ID:        uuid.New(),
		TeamID:    teamID,
		Code:      code,
		CreatedAt: time.Now(),
	}, nil
}

// generateInviteCode creates a unique alphanumeric code (e.g., "ABC123XYZ")
func generateInviteCode() (string, error) {
	// Generate 6 random bytes
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// Encode to base32 and clean up
	code := base32.StdEncoding.EncodeToString(b)
	code = strings.ReplaceAll(code, "=", "")
	code = strings.ToUpper(code)

	// Take first 9 characters for readability
	if len(code) > 9 {
		code = code[:9]
	}

	return code, nil
}

