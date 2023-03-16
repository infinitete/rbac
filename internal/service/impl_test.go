package service

import (
	"github.com/infinitete/rbac"
	"github.com/infinitete/rbac/internal/io"
	"github.com/infinitete/rbac/internal/model"
	"testing"
)

var resourceStore rbac.Store[model.Resource]
var userStore rbac.Store[model.User]
var roleStore rbac.Store[model.Role]
var permStore rbac.Store[model.Perm]

var resourceService *ResourceService
var userService *UserService
var roleService *RoleService
var permService *PermService

func init() {
	resourceStore = io.GetStore[*model.Resource]()
	userStore = io.GetStore[*model.User]()
	roleStore = io.GetStore[*model.Role]()
	permStore = io.GetStore[*model.Perm]()

	resourceService = GetResourceService(resourceStore)
	userService = GetUserService(userStore)
	roleService = GetRoleService(roleStore)
	permService = GetPermService(permStore)
}

func TestRbac(t *testing.T) {
	resource := &model.Resource{UniqueKey: "resource"}
	user := &model.User{
		Name:  "test",
		Roles: []string{},
	}
	role := &model.Role{
		UniqueKey: "test_role",
		Perms:     []string{},
	}

	perm := &model.Perm{
		UniqueKey:   "test_perm",
		Description: "it just for test",
	}

	_ = resourceService.Create(resource)
	_ = userService.Create(user)
	_ = roleService.Create(role)
	_ = permService.Create(perm)

	err := resource.Access(perm)
	t.Logf("AccessResult: %v", err) // access denied

	_ = role.AddPerm(perm)

	_ = resource.AddPerm(perm)
	err = resource.Access(perm)
	t.Logf("AccessResult: %v", err)

	user.AddRole(role)
	var user1 *model.User
	_ = userStore.Update(user)
	user1, err = userStore.Read(user.Key())

	t.Logf("Roles: %p - %p", user1, user)
}
