package usersHandlers

import (
	"github.com/NIRVXSH/NIRVXSH-shop-project/config"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/users/usersUsecases"
)

type IUsersHandler interface {
}

type usersHandler struct {
	cfg          config.IConfig
	usersUsecase usersUsecases.IUsersUsecase
}

func UsersHandler(cfg config.IConfig, usersUsecase usersUsecases.IUsersUsecase) IUsersHandler {
	return &usersHandler{
		cfg:          cfg,
		usersUsecase: usersUsecase,
	}
}
