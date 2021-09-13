package createuser

import (
	"github.com/raismaulana/assetsP/domain/repository"
	"github.com/raismaulana/assetsP/domain/service"
)

// Outport of CreateUser
type Outport interface {
	repository.FindUserByEmailRepo
	repository.ReadOnlyDB
	repository.SaveUserRepo
	repository.TransactionDB
	service.HashPasswordService
}
