package resource

import (
	. "github.com/talbright/go-desk/types"
)

type User struct {
	Name           *string    `json:"name,omitempty"`
	PublicName     *string    `json:"public_name,omitempty"`
	Email          *string    `json:"email,omitempty"`
	EmailVerified  *bool      `json:"email_verified,omitempty"`
	Available      *bool      `json:"available,omitempty"`
	Avatar         *string    `json:"avatar,omitempty"`
	Level          *string    `json:"level,omitempty"`
	CreatedAt      *Timestamp `json:"created_at,omitempty"`
	UpdatedAt      *Timestamp `json:"updated_at,omitempty"`
	CurrentLoginAt *Timestamp `json:"current_login_at,omitempty"`
	LastLoginAt    *Timestamp `json:"last_login_at,omitempty"`
	Resource
}

func NewUser() *User {
	user := &User{}
	user.InitializeResource(user)
	return user
}

func (c User) String() string {
	return Stringify(c)
}
