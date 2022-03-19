package dto

import "time"

type User struct {
	ID        uint64
	Balance   uint64    `gorm:"default:0;comment:餘額"`
	CreatedAt time.Time `gorm:"comment:創建時間"`
	UpdatedAt time.Time `gorm:"comment:更新時間"`
}

func (User) TableName() string {
	return "user"
}
