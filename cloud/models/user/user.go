package user

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/mzz2017/grip/cloud/common/crypto"
	"github.com/mzz2017/grip/cloud/dao"
)

type User struct {
	gorm.Model
	Sub          string `gorm:"not null;unique"`
	Password     string `gorm:"not null"`
	CompanyID    int    `gorm:"not null"`
	Department   string `gorm:"not null"`
	PhoneNumber  string
	WechatID     string
	EnterpriseID string
	Name         string `gorm:"not null"`
	Admin        bool   `gorm:"not null"`
}

func (user *User) Insert() error {
	user.Password = crypto.CryptoPwd(user.Password)
	return dao.DB().Table("users").Create(user).Error
}

func (user *User) Count() (cnt int64) {
	dao.DB().Table("users").Where(user).Count(&cnt)
	return
}

func (user User) Find() (u User, err error) {
	err = dao.DB().Table("users").Where(&user).First(&u).Error
	return
}
func (user User) Update() (err error) {
	var eu User
	if dao.DB().Model(&user).First(&eu); eu.CreatedAt.IsZero() {
		return errors.New("invalid ID")
	}
	if user.Password != "$__hidden_PASSWORD" {
		user.Password = crypto.CryptoPwd(user.Password)
	} else {
		user.Password = eu.Password
	}
	return dao.DB().Model(&user).Updates(user).Error
}

func (user User) Delete() error {
	return dao.DB().Unscoped().Delete(&user).Error
}
