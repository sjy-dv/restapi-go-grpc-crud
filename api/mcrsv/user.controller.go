package mcrsv

import (
	"time"

	"github.com/jinzhu/gorm"
)


type User struct {
	Idx      int    `gorm:"primary_key;auto_increment;not null" json:"idx"`
	Username string `json:"username"`
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	CreatedAt    time.Time `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updatedAt" json:"updated_at"`
}

func (u *User) CreateUser(db *gorm.DB) error {

	err := db.Debug().Create(&u).Error

	if err != nil {
		return err
	}
	return nil
}