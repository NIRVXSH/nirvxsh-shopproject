package servers

import (
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/middlewares/middlewaresHandlers"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/middlewares/middlewaresRepositories"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/middlewares/middlewaresUsecases"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/monitor/monitorHandlers"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/users/usersHandlers"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/users/usersRepositories"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/users/usersUsecases"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
	UsersModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewaresHandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewaresHandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewaresHandlers.IMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewaresHandlers.MiddlewaresHandler(s.cfg, usecase)

}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}

func (m *moduleFactory) UsersModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)

	router := m.r.Group("/users")

	router.Post("/signup", handler.SignUpCustomer)

}
