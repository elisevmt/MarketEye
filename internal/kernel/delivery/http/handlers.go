package http

import (
	"MarketEye/config"
	"MarketEye/internal/kernel"
	"MarketEye/internal/models"
	"MarketEye/pkg/logger"
	"MarketEye/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type kernelHandlers struct {
	cfg      *config.Config
	logger   *logger.ApiLogger
	kernelUC kernel.UseCase
}

func NewKernelHandlers(cfg *config.Config, log *logger.ApiLogger, kernelUS kernel.UseCase) kernel.Handlers {
	return &kernelHandlers{cfg: cfg, logger: log, kernelUC: kernelUS}
}

func (p *kernelHandlers) Price() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.FetchPricesParams
		if err := utils.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(p.kernelUC.FetchPrices(params))
	}
}

func (p *kernelHandlers) Depth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params models.FetchOrderBookParams
		if err := utils.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(p.kernelUC.FetchOrderBook(params))
	}
}

func (p *kernelHandlers) FetchMarketList() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(p.kernelUC.FetchMarketList())
	}
}
