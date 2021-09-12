package rdbms

import (
	"context"

	"github.com/raismaulana/assetsP/infrastructure/database"
	"github.com/raismaulana/assetsP/infrastructure/envconfig"
	"github.com/raismaulana/assetsP/infrastructure/migration"
	"gorm.io/gorm"
)

type RDBMSGateway struct {
	database.GormReadOnlyImpl
	database.GormTransactionImpl
}

// NewRDBMSGateway ...
func NewRDBMSGateway(ctx context.Context, env *envconfig.EnvConfig, db *gorm.DB) (*RDBMSGateway, error) {
	err := migration.RDBMSMigration(ctx, db, env)
	if err != nil {
		return nil, err
	}

	return &RDBMSGateway{
		GormReadOnlyImpl: database.GormReadOnlyImpl{
			DB: db,
		},
		GormTransactionImpl: database.GormTransactionImpl{
			DB: db,
		},
	}, nil
}
