package kernel

import "github.com/gofiber/fiber/v2"

type Handlers interface {
	Price() fiber.Handler
	Depth() fiber.Handler
	FetchMarketList() fiber.Handler
}
