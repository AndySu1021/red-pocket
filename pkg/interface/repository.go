package iface

import (
	"context"
	"database/sql"
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
	Get(ctx context.Context, tx *gorm.DB, data interface{}, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (int64, error)
	GetOne(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Create(ctx context.Context, tx *gorm.DB, data interface{}, scopes ...func(*gorm.DB) *gorm.DB) error
	Update(ctx context.Context, tx *gorm.DB, opt WhereOption, col UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error
	Delete(ctx context.Context, tx *gorm.DB, model Model, opt WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error
	Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
}
