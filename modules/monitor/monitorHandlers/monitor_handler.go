package monitorHandlers

import (
	"github.com/NIRVXSH/NIRVXSH-shop-project/config"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/entities"
	"github.com/NIRVXSH/NIRVXSH-shop-project/modules/monitor"
	"github.com/gofiber/fiber/v2"
)

type IMontitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMontitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
