package dto

import "time"

type RedPacket struct {
	ID         uint64
	ActivityID uint64    `gorm:"comment:活動 ID"`
	UserID     uint64    `gorm:"comment:用戶 ID"`
	Amount     uint64    `gorm:"comment:紅包金額"`
	CreatedAt  time.Time `gorm:"comment:創建時間"`
	UpdatedAt  time.Time `gorm:"comment:更新時間"`
}

func (RedPacket) TableName() string {
	return "red_packet"
}
