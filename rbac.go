package rbac

import "errors"

var AccessDenied = errors.New("access denied")
var DelPermErr = errors.New("del perm error")
var AddPermErr = errors.New("add perm error")

type Storable interface {
	Key() string
}

type User interface {
	Key() string
	HasRole(role Role) bool
	AddRole(role Role) error
	GetRoles() []string
}

type Role interface {
	Key() string
	AddPerm(perm Perm) error
	DelPerm(perm Perm) error
	HasPerm(perm Perm) bool
	GetPerms() []string
}

type Perm interface {
	Key() string
}

type Resource interface {
	Key() string
	DelPerm(Perm) error
	AddPerm(Perm) error
	Access(Perm) error
}

type Store[T any] interface {
	Read(key string) (*T, error)
	Create(model *T) error
	Update(model *T) error
	Delete(key string) error
}

type Service[T any] interface {
	Store[T]
	SetStore(store Store[T])
}

type UserService Service[User]
type RoleService Service[Role]
type PermService Service[Perm]
type ResourceService Service[Resource]
