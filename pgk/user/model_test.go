package user

import (
	"github.com/stretchr/testify/assert"
	"github.com/tonsV2/todo-go/pgk/group"
	"testing"
)

func TestIsMemberOfByIdNegative(t *testing.T) {
	group0 := &group.Group{}
	group0.ID = 0
	group1 := &group.Group{}
	group1.ID = 1
	user := &User{Groups: []group.Group{*group0, *group1}}

	isMember := user.IsMemberOfById(2)

	assert.False(t, isMember)
}

func TestIsMemberOfByIdPositive(t *testing.T) {
	group0 := &group.Group{}
	group0.ID = 0
	group1 := &group.Group{}
	group1.ID = 1
	user := &User{Groups: []group.Group{*group0, *group1}}

	isMember := user.IsMemberOfById(1)

	assert.True(t, isMember)
}

func TestUserIsAdministratorNegative(t *testing.T) {
	group0 := &group.Group{Name: "something0"}
	group1 := &group.Group{Name: "something1"}
	user := &User{Groups: []group.Group{*group0, *group1}}

	isAdministrator := user.IsAdministrator()

	assert.False(t, isAdministrator)
}

func TestUserIsAdministratorPositive(t *testing.T) {
	group0 := &group.Group{Name: "something"}
	group1 := &group.Group{Name: group.AdministratorGroupName}
	user := &User{Groups: []group.Group{*group0, *group1}}

	isAdministrator := user.IsAdministrator()

	assert.True(t, isAdministrator)
}
func TestUserIsMemberOfNegative(t *testing.T) {
	group0 := &group.Group{Name: "something0"}
	group1 := &group.Group{Name: "something1"}
	user := &User{Groups: []group.Group{*group0, *group1}}

	isMember := user.IsMemberOf("something2")

	assert.False(t, isMember)
}

func TestUserIsMemberOfPositive(t *testing.T) {
	group0 := &group.Group{Name: "something0"}
	group1 := &group.Group{Name: "something1"}
	user := &User{Groups: []group.Group{*group0, *group1}}

	isMember := user.IsMemberOf("something1")

	assert.True(t, isMember)
}
