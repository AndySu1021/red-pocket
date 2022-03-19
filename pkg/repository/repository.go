package repository

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"gorm.io/gorm"
	iface "red-packet/pkg/interface"
)

type repository struct {
	db *gorm.DB
}

var Module = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(
		Migration,
	),
)

type Params struct {
	fx.In

	DB *gorm.DB
}

func New(p Params) (iface.IRepository, error) {
	repo := &repository{
		db: p.DB,
	}
	return repo, nil
}

func (repo *repository) GetDB() *gorm.DB {
	return repo.db
}

func (repo *repository) Get(ctx context.Context, tx *gorm.DB, data interface{}, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) (total int64, err error) {
	if tx == nil {
		tx = repo.db.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)

	db := tx.Table(opt.TableName()).Model(data).Scopes(opt.Where)

	db = opt.Preload(tx)
	err = db.Scopes(opt.Where).Find(data).Error
	if err != nil {
		return
	}

	return
}

func (repo *repository) GetOne(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.db.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	tx = opt.Preload(tx)
	err := tx.Table(model.TableName()).Scopes(opt.Where).First(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Create(ctx context.Context, tx *gorm.DB, data interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.db.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	err := tx.Create(data).Error
	return err
}

func (repo *repository) Update(ctx context.Context, tx *gorm.DB, opt iface.WhereOption, col iface.UpdateColumns, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.db.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	if opt.IsEmptyWhereOpt() {
		return errors.New("database: Update err: where condition can't empty")
	}
	err := tx.Table(opt.TableName()).Scopes(opt.Where).Updates(col.Columns()).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) Delete(ctx context.Context, tx *gorm.DB, model iface.Model, opt iface.WhereOption, scopes ...func(*gorm.DB) *gorm.DB) error {
	if tx == nil {
		tx = repo.db.WithContext(ctx)
	}
	tx = tx.Scopes(scopes...)
	err := tx.Scopes(opt.Where).Delete(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Transaction(ctx context.Context, fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)
		return fc(tx)
	}, opts...)
}
