package entity

import "errors"

var (
	ErrTeamNameRequired   = errors.New("team name is required")
	ErrRoleNameRequired   = errors.New("role name is required")
	ErrMemberNameRequired = errors.New("member name is required")
	ErrMemberAlreadyExists = errors.New("member already exists in this team")
)

