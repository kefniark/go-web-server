package core

import (
	dbClient "github.com/kefniark/go-web-server/gen/db"
	"github.com/rs/zerolog"
)

type ServerOptions struct {
	Logger *zerolog.Logger
	DB     *dbClient.Queries
}

// Result of the Auth Middleware
// If the user is known, pass it to the rest of the system through context

type AuthInfo interface {
	Token() string
	Name() string
}

type AuthUserInfo struct {
	token string
	name  string
}

func (info AuthUserInfo) Token() string {
	return info.token
}

func (info AuthUserInfo) Name() string {
	return info.name
}

func NewUserInfo(token string, name string) AuthInfo {
	return &AuthUserInfo{
		token: token,
		name:  name,
	}
}