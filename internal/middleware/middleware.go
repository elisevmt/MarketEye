package middleware

import (
	"MarketEye/config"
	"github.com/gofiber/fiber/v2"
)

type MDWManager struct {
	cfg *config.Config
}

func NewMDWManager(cfg *config.Config) *MDWManager {
	return &MDWManager{
		cfg: cfg,
	}
}

func (mw *MDWManager) TerminalMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if c.Get("APIKey") != mw.cfg.TerminalAccess.APIKey {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
