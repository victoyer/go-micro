// @Time    : 2019-06-14 20:24
// @Author  : victor！！
// @FileName: extension.go
// @Software: IntelliJ IDEA

package userPKG

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (this *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("Id", uuid.NewV4().String())
}
