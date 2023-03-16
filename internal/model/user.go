package model

import "github.com/infinitete/rbac"

type User struct {
	Name  string
	Roles []string
}

func (user *User) Key() string {
	return user.Name
}

func (user *User) HasRole(role rbac.Role) bool {
	for _, r := range user.Roles {
		if role.Key() == r {
			return true
		}
	}

	return false
}

func (user *User) AddRole(role rbac.Role) error {
	if user.HasRole(role) {
		return nil
	}
	user.Roles = append(user.Roles, role.Key())
	return nil
}

func (user *User) GetRoles() []string {
	return user.Roles
}
