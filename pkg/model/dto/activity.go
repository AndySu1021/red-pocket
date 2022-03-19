package dto

import "time"

type Activity struct {
	ID        uint64
	Count     uint64    `gorm:"comment:紅包總數"`
	Amount    uint64    `gorm:"comment:紅包總金額"`
	CreatedBy uint64    `gorm:"comment:創建使用者"`
	CreatedAt time.Time `gorm:"comment:創建時間"`
	UpdatedAt time.Time `gorm:"comment:更新時間"`
}

func (Activity) TableName() string {
	return "activity"
}
