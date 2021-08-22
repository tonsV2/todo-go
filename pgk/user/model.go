package user

import (
	"github.com/tonsV2/todo-go/pgk/group"
	"gorm.io/gorm"
	"sort"
)

type User struct {
	gorm.Model
	Email    string        `gorm:"unique;"`
	Password string        `json:"-"`
	Groups   []group.Group `gorm:"many2many:user_groups;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) IsMemberOfById(groupId uint) bool {
	groups := u.Groups

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].ID <= groups[j].ID
	})

	index := sort.Search(len(groups), func(i int) bool {
		return groups[i].ID >= groupId
	})

	return index < len(groups) && groups[index].ID == groupId
}

func (u *User) IsMemberOf(group string) bool {
	groups := u.Groups

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Name <= groups[j].Name
	})

	index := sort.Search(len(groups), func(i int) bool {
		return groups[i].Name >= group
	})

	return index < len(groups) && groups[index].Name == group
}

func (u *User) IsAdministrator() bool {
	return u.IsMemberOf(group.AdministratorGroupName)
}
