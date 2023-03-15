package internal

import "github.com/infinitete/rbac"

type User struct {
}

func (user *User) HasRole(role rbac.Role) bool {
	panic("TODO")
}

func (user *User) AddRole(role rbac.Role) error {
	panic("TODO")
}

func (user *User) GetRoles() []rbac.Role {
	panic("TODO")
}

func (user *User) HasPerm(perm Perm) bool {
	panic("TODO")
}

func (user *User) AddPerms(perms ...Perm) error {
	panic("TODO")
}
