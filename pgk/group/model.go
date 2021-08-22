package group

import (
	"github.com/tonsV2/todo-go/pgk/user"
	"gorm.io/gorm"
)

const AdministratorGroupName = "administrators"

type Group struct {
	gorm.Model
	Name     string      `gorm:"unique;"`
	Hostname string      `gorm:"unique;"`
	Users    []user.User `gorm:"many2many:user_groups;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
