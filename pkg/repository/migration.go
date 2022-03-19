package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	iface "red-packet/pkg/interface"
	"red-packet/pkg/model/dto"
)

func Migration(repo iface.IRepository) error {
	conn := repo.GetDB()
	conn.DisableForeignKeyConstraintWhenMigrating = true
	_conn := conn.Session(
		&gorm.Session{
			Logger: logger.Default.LogMode(logger.Warn),
		},
	)
	err := _conn.AutoMigrate(
		&dto.User{},
		&dto.RedPacket{},
		&dto.Activity{},
	)
	if err != nil {
		return err
	}

	return nil
}
