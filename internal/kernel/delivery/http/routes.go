package http

import (
	"MarketEye/internal/kernel"
	"MarketEye/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func MapKernelRoutes(kernelRoutes fiber.Router, h kernel.Handlers, mw *middleware.MDWManager) {
	kernelRoutes.Post("/price", mw.TerminalMiddleware(), h.Price())
	kernelRoutes.Post("/depth", mw.TerminalMiddleware(), h.Depth())
	kernelRoutes.Get("/market/list", mw.TerminalMiddleware(), h.FetchMarketList())
}
