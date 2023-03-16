package model

import "github.com/infinitete/rbac"

type Role struct {
	UniqueKey string
	Perms     []string
}

func (role *Role) Key() string {
	return role.UniqueKey
}

func (role *Role) AddPerm(perm rbac.Perm) error {
	if role.HasPerm(perm) {
		return nil
	}

	role.Perms = append(role.Perms, perm.Key())
	return nil
}

func (role *Role) DelPerm(perm rbac.Perm) error {
	if !role.HasPerm(perm) {
		return nil
	}

	var nextPerms []string
	for _, p := range role.Perms {
		if p == perm.Key() {
			continue
		}
		nextPerms = append(nextPerms, p)
	}
	role.Perms = nextPerms

	return nil
}

func (role *Role) HasPerm(perm rbac.Perm) bool {
	for _, p := range role.Perms {
		if p == perm.Key() {
			return true
		}
	}

	return false
}

func (role *Role) GetPerms() []string {
	return role.Perms
}
