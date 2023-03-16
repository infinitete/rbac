package service

import (
	"github.com/infinitete/rbac"
)

type serviceImpl[T any] struct {
	store rbac.Store[T]
}

func (u *serviceImpl[T]) SetStore(store rbac.Store[T]) {
	u.store = store
}

func (u *serviceImpl[T]) Create(key string, model T) (T, error) {
	return u.store.Create(key, model)
}

func (u *serviceImpl[T]) Read(key string) (T, error) {
	return u.store.Read(key)
}

func (u *serviceImpl[T]) Update(key string, model T) (T, error) {
	return u.store.Update(key, model)
}

func (u *serviceImpl[T]) Delete(key string) error {
	return u.store.Delete(key)
}

type RoleService struct {
	serviceImpl[rbac.Role]
}

type UserService struct {
	serviceImpl[rbac.User]
}

type PermService struct {
	serviceImpl[rbac.Perm]
}

type ResourceService struct {
	serviceImpl[rbac.Resource]
}

func GetRoleService(store rbac.Store[rbac.Role]) *RoleService {
	s := &RoleService{}
	s.SetStore(store)
	return s
}

func GetUserService(store rbac.Store[rbac.User]) *UserService {
	s := &UserService{}
	s.SetStore(store)
	return s
}

func GetPermService(store rbac.Store[rbac.Perm]) *PermService {
	s := &PermService{}
	s.SetStore(store)
	return s
}

func GetResourceService(store rbac.Store[rbac.Resource]) *ResourceService {
	s := ResourceService{}
	s.SetStore(store)
	return &s
}
