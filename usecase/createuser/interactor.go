package createuser

import (
	"context"

	"github.com/raismaulana/assetsP/application/apperror"
	"github.com/raismaulana/assetsP/domain/entity"
	"github.com/raismaulana/assetsP/domain/repository"
	"github.com/raismaulana/assetsP/infrastructure/log"
)

//go:generate mockery --name Outport -output mocks/

type createUserInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase CreateUser
func NewUsecase(outputPort Outport) Inport {
	return &createUserInteractor{
		outport: outputPort,
	}
}

// Execute the usecase CreateUser
func (r *createUserInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}
	err := repository.ReadOnly(ctx, r.outport, func(ctx context.Context) error {
		userObj, err := r.outport.FindUserByEmail(ctx, req.Email)
		if userObj != nil {
			log.Error(ctx, err.Error())
			return apperror.EmailAlreadyExist
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = repository.WithTransaction(ctx, r.outport, func(ctx context.Context) error {
		req.Password, err = r.outport.HashPassword(ctx, req.Password)
		if err != nil {
			return err
		}
		userObj, err := entity.NewUser(entity.UserRequest{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			Role:     "peasant",
		})
		if err != nil {
			return err
		}
		err = r.outport.SaveUser(ctx, userObj)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
