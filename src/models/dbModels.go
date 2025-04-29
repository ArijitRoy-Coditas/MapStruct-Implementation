package models

import "gorm.io/gorm"

type DatabaseConfiguration struct {
	GormDB *gorm.DB
}

type User struct {
	Id          uint32 `gorm:"column:id;primaryKey;index;autoIncrement;" json:"id"`
	Username    string `gorm:"column:username;uniqueIndex" json:"username"`
	Name        string `gorm:"column:name" json:"name"`
	Email       string `gorm:"column:email;uniqueIndex" json:"email"`
	Password    string `gorm:"column:password" json:"password"`
	PanCard     string `gorm:"column:panCard;uniqueIndex" json:"panCard"`
	PhoneNumber uint64 `gorm:"column:phoneNumber" json:"phoneNumber"`
}
