package service

import (
	"github.com/infinitete/rbac"
	"github.com/infinitete/rbac/internal/io"
	"github.com/infinitete/rbac/internal/model"
	"testing"
)

var resourceStore rbac.Store[rbac.Resource]
var userStore rbac.Store[rbac.User]
var roleStore rbac.Store[rbac.Role]
var permStore rbac.Store[rbac.Perm]

var resourceService *ResourceService
var userService *UserService
var roleService *RoleService
var permService *PermService

func init() {
	resourceStore = io.GetStore[rbac.Resource]()
	userStore = io.GetStore[rbac.User]()
	roleStore = io.GetStore[rbac.Role]()
	permStore = io.GetStore[rbac.Perm]()

	resourceService = GetResourceService(resourceStore)
	userService = GetUserService(userStore)
	roleService = GetRoleService(roleStore)
	permService = GetPermService(permStore)
}

func TestRbac(t *testing.T) {
	resource := &model.Resource{UniqueKey: "resource"}
	user := model.User{
		Name:  "test",
		Roles: []string{},
	}
	role := model.Role{
		UniqueKey: "test_role",
		Perms:     []string{},
	}
	perm := model.Perm{
		UniqueKey:   "test_perm",
		Description: "it just for test",
	}

	r1, _ := resourceService.Create("", resource)
	_ = r1.(*model.Resource).Perms
	resource, _ = resourceService.Create(resource.Key(), resource)
	_ = userService.Create(user.Key(), user)
	_ = roleService.Create(role.Key(), role)
	_ = permService.Create(perm.Key(), perm)

	err := resource.Access(perm)
	t.Logf("AccessResult: %v", err) // access denied

	_ = role.AddPerm(perm)

	_ = resource.AddPerm(perm)
	err = resource.Access(perm)
	t.Logf("AccessResult: %v", err)

	user.AddRole(role)
	user1 := model.User{}
	user, _ = userStore.Update(user.Key(), user)
	user1, err = userStore.Read(user.Key())

	t.Logf("Roles: %v - %v", user.Roles, user1.Roles)
}
