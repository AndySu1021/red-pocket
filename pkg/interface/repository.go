package iface

import (
	"context"
	"gorm.io/gorm"
)

type Model interface {
	TableName() string
}

type WhereOption interface {
	Model
	Where(db *gorm.DB) *gorm.DB
	Preload(db *gorm.DB) *gorm.DB
	IsEmptyWhereOpt() bool
}

type UpdateColumns interface {
	Columns() interface{}
}

type IRepository interface {
	GetDB() *gorm.DB
	Get(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Create(ctx context.Context, tx *gorm.DB, data Model, scopes ...func(*gorm.DB) *gorm.DB) error
	Update(ctx context.Context, tx *gorm.DB, opt WhereOption, col UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error
	Delete(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
}
