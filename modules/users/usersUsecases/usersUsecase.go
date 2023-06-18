package usersUsecases

import (
	"github.com/NIRVXSH/NIRVXSH-shop-project/config"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/users/usersRepositories"
)

type IUsersUsecase interface {
}

type usersUsecase struct {
	cfg             config.IConfig
	usersRepository usersRepositories.IUsersRepository
}

func UsersUsecase(cfg config.IConfig, usersRepository usersRepositories.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		cfg:             cfg,
		usersRepository: usersRepository,
	}
}
