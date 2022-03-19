package option

import (
	"demo/pkg/model/dto"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type UserOption struct {
	User dto.User
}

func (where *UserOption) Where(db *gorm.DB) *gorm.DB {
	db = db.Where(where.User)

	return db
}

func (where *UserOption) IsEmptyWhereOpt() bool {
	return reflect.DeepEqual(where.User, dto.User{})
}

func (where *UserOption) TableName() string {
	return where.User.TableName()
}

func (where *UserOption) Preload(db *gorm.DB) *gorm.DB {
	return db
}

type UserUpdateColumn struct {
	Balance   uint64
	UpdatedAt time.Time
}

func (col *UserUpdateColumn) Columns() interface{} {
	return col
}
