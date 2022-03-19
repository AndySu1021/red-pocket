package option

import (
	"gorm.io/gorm"
	"red-packet/pkg/model/dto"
	"reflect"
	"time"
)

type RedPacketOption struct {
	RedPacket dto.RedPacket
}

func (where *RedPacketOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.RedPacket)
	return db
}

func (where *RedPacketOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

func (where *RedPacketOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.RedPacket, dto.RedPacket{})
}

func (where *RedPacketOption) TableName() string {
	return where.RedPacket.TableName()
}

type RedPacketUpdateColumn struct {
	UserID    uint64
	Status    int8
	UpdatedAt time.Time
}

func (col *RedPacketUpdateColumn) Columns() interface{} {
	return col
}
