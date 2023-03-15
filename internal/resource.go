package internal

import "github.com/infinitete/rbac"

type Resource struct {
	Perms []string
}

func (r *Resource) DelPerm(perm rbac.Perm) error {
	nextPerms := []string{}
	for _, p := range r.Perms {
		if perm.Key() != p {
			nextPerms = append(nextPerms, p)
		}
	}
	r.Perms = nextPerms

	return nil
}

func (r *Resource) AddPerms(perms ...rbac.Perm) error {
	permsKeyMap := make(map[string]struct{}, len(r.Perms))
	if len(r.Perms) > 0 {
		for _, key := range r.Perms {
			permsKeyMap[key] = struct{}{}
		}
	}

	nextPerms := []string{}
	for _, perm := range perms {
		if _, ok := permsKeyMap[perm.Key()]; !ok {
			nextPerms = append(nextPerms, perm.Key())
		}
	}

	r.Perms = append(r.Perms, nextPerms...)

	return nil
}

func (r *Resource) TryAccessWithUser(user rbac.User) error {
	for _, role := range user.GetRoles() {
		if r.TryAccessWithRole(role) == nil {
			return nil
		}
	}

	return rbac.AccessDenied
}

func (r *Resource) TryAccessWithRole(role rbac.Role) error {
	for _, perm := range role.GetPerms() {
		if r.TryAccessWithPerm(perm) == nil {
			return nil
		}
	}

	return rbac.AccessDenied
}

func (r *Resource) TryAccessWithPerm(perm rbac.Perm) error {
	for _, p := range r.Perms {
		if perm.Key() == p {
			return nil
		}
	}

	return rbac.AccessDenied
}
