package repository

import (
	"context"

	"github.com/raismaulana/assetsP/domain/entity"
)

type SaveUserRepo interface {
	SaveUser(ctx context.Context, obj *entity.User) error
}

type FindUserByEmailRepo interface {
	FindUserByEmail(ctx context.Context, someID string) ([]*entity.User, error)
}
