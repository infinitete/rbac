package model

import "github.com/infinitete/rbac"

type Resource struct {
	UniqueKey string
	Perms     []string
}

func (r *Resource) Key() string {
	return r.UniqueKey
}

func (r *Resource) DelPerm(perm rbac.Perm) error {
	var nextPerms []string
	for _, p := range r.Perms {
		if perm.Key() != p {
			nextPerms = append(nextPerms, p)
		}
	}
	r.Perms = nextPerms

	return nil
}

func (r *Resource) AddPerm(perm rbac.Perm) error {
	for _, p := range r.Perms {
		if p == perm.Key() {
			return nil
		}
	}
	r.Perms = append(r.Perms, perm.Key())
	return nil
}

func (r *Resource) Access(perm rbac.Perm) error {
	for _, p := range r.Perms {
		if perm.Key() == p {
			return nil
		}
	}

	return rbac.AccessDenied
}
