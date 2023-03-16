package service

import (
	"github.com/infinitete/rbac"
	"github.com/infinitete/rbac/internal/model"
)

type serviceImpl[T any] struct {
	store rbac.Store[T]
}

func (u *serviceImpl[T]) SetStore(store rbac.Store[T]) {
	u.store = store
}

func (u *serviceImpl[T]) Create(model *T) error {
	return u.store.Create(model)
}

func (u *serviceImpl[T]) Read(key string) (*T, error) {
	return u.store.Read(key)
}

func (u *serviceImpl[T]) Update(model *T) error {
	return u.store.Update(model)
}

func (u *serviceImpl[T]) Delete(key string) error {
	return u.store.Delete(key)
}

type RoleService struct {
	serviceImpl[model.Role]
}

type UserService struct {
	serviceImpl[model.User]
}

type PermService struct {
	serviceImpl[model.Perm]
}

type ResourceService struct {
	serviceImpl[model.Resource]
}

func GetRoleService(store rbac.Store[model.Role]) *RoleService {
	s := &RoleService{}
	s.SetStore(store)
	return s
}

func GetUserService(store rbac.Store[model.User]) *UserService {
	s := &UserService{}
	s.SetStore(store)
	return s
}

func GetPermService(store rbac.Store[model.Perm]) *PermService {
	s := &PermService{}
	s.SetStore(store)
	return s
}

func GetResourceService(store rbac.Store[model.Resource]) *ResourceService {
	s := ResourceService{}
	s.SetStore(store)
	return &s
}
