package rdbms

import (
	"context"

	"github.com/raismaulana/assetsP/application/apperror"
	"github.com/raismaulana/assetsP/domain/entity"
	"github.com/raismaulana/assetsP/infrastructure/database"
	"github.com/raismaulana/assetsP/infrastructure/log"
)

func (r *RDBMSGateway) SaveUser(ctx context.Context, obj *entity.User) error {
	log.Info(ctx, "called")
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return apperror.ERR500.Var(err)
	}
	err = db.Save(obj).Error
	if err != nil {
		log.Error(ctx, err.Error())
		return apperror.ERR400.Var(err)
	}
	return nil
}
