package rbac

import "errors"

var AccessDenied = errors.New("access denied")
var DelPermErr = errors.New("del perm error")
var AddPermErr = errors.New("add perm error")

type User interface {
	HasRole(role Role) bool
	AddRole(role Role) error
	GetRoles() []Role

	HasPerm(perm Perm) bool
	AddPerm(perm Perm) error
}

type Role interface {
	Key() string
	HasPerm(perm Perm) bool
	AddPerm(perm Perm) error
	DelPerm(perm Perm) error

	GetPerms() []Perm
}

type Perm interface {
	Key() string
}

type Resource interface {
	Key() string
	DelPerm(perm Perm) error
	AddPerm(perm Perm) error
	TryAccessWithUser(user User) error
	TryAccessWithRole(role Role) error
	TryAccessWithPerm(perm Perm) error
}

type Reader[T any] func(model T) error
type Writer[T any] func(key string) (T, error)

type RbacService[T any] interface {
	SetReader(reader Reader[T])
	SetWriter(writer Writer[T])
}

type UserService RbacService[User]
type RoleService RbacService[Role]
type PermService RbacService[Perm]
type ResourceService RbacService[Resource]

// type UserWriter Writer[User]
// type RoleWriter Writer[Role]
// type PermWriter Writer[Perm]

// type UserReader Reader[User]
// type RoleReader Reader[Role]
// type PermReader Reader[Perm]
