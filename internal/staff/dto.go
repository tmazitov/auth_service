package staff

import (
	"time"
)

type AuthMethod int

const (
	EmailAuthMethod  AuthMethod = 1
	GoogleAuthMethod AuthMethod = 2
	GithubAuthMethod AuthMethod = 3
)

type TokenPair struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}

type UserAuthMethod struct {
	Id           int        `bun:"id,pk,autoincrement"`
	UserId       int        `bun:"user_id,pk"`
	AuthMethodId AuthMethod `bun:"auth_method_id,pk"`
	LastAuthAt   *time.Time `bun:"last_auth_at, default:current_timestamp"`
}

type UserAuth struct {
	Id         int        `bun:"id,pk,autoincrement"`
	FirstName  string     `bun:"first_name"`
	LastName   string     `bun:"last_name"`
	Email      string     `bun:"email,unique"`
	LastAuthAt *time.Time `bun:"last_auth_at, default:current_timestamp"`
	CreatedAt  *time.Time `bun:"created_at, default:current_timestamp"`
	UpdatedAt  *time.Time `bun:"updated_at, default:current_timestamp"`
}
