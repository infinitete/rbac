package internal

import "github.com/infinitete/rbac"

type Role struct {
	Key   string
	Perms []string
}

func (role *Role) HasPerm(perm rbac.Perm) bool {
	for _, p := range role.Perms {
		if p == perm.Key() {
			return true
		}
	}

	return false
}
func (role *Role) AddPerms(perm rbac.Perm) error {
	if role.HasPerm(perm) {
		return nil
	}

	role.Perms = append(role.Perms, perm.Key())

	return nil
}
func (role *Role) DelPerms(perm rbac.Perm) error {
	nextPerms := []string{}
	for _, p := range role.Perms {
		if p == perm.Key() {
			continue
		}
		nextPerms = append(nextPerms, p)
	}
	return nil
}

func (role *Role) GetPerms() []rbac.Perm {
	panic("TODO")
}
