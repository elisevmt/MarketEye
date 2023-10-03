package server

import (
//	bestchangeUseCase "MarketEye/internal/bestchangeTree/usecase"
	binanceUseCase "MarketEye/internal/binanceTree/usecase"
//	garantexUseCase "MarketEye/internal/garantexTree/usecase"
	kernelHttp "MarketEye/internal/kernel/delivery/http"
	kernelUseCase "MarketEye/internal/kernel/usecase"
	apiMiddlewares "MarketEye/internal/middleware"
	"MarketEye/pkg/fhttp"
	"MarketEye/pkg/graphZero"
	"MarketEye/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	server_logger "github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *Server) MapHandlers(app *fiber.App, logger *logger.ApiLogger) error {
	fhttpClient := fhttp.NewClient(s.cfg)

	marketMap := make(map[string]*graphZero.Market)

	binanceMarket := graphZero.NewMarket("binance")
	marketMap["binance"] = binanceMarket

	//garantexMarket := graphZero.NewMarket("garantex")
	//marketMap["garantex"] = garantexMarket

	//bestchangeMarket := graphZero.NewMarket("bestchange")
	//marketMap["bestchange"] = bestchangeMarket

	rootNode := graphZero.NewNode(binanceMarket, "USDT")

	kernelTree := graphZero.NewTree(rootNode)

	//garantexTreeUC := garantexUseCase.NewGarantexTreeUC(s.cfg, garantexMarket, kernelTree, fhttpClient, logger)
	binanceTreeUC := binanceUseCase.NewBinanceTreeUC(s.cfg, binanceMarket, kernelTree, fhttpClient, logger)
	//bestchangeTreeUC := bestchangeUseCase.NewBestchangeTreeUC(s.cfg, bestchangeMarket, kernelTree, fhttpClient, logger)

	go binanceTreeUC.TreeLoad()
	//go garantexTreeUC.TreeLoad()
	//go bestchangeTreeUC.TreeLoad()

	kernelUC := kernelUseCase.NewKernelUseCase(s.cfg, kernelTree, &marketMap)

	kernelHandlers := kernelHttp.NewKernelHandlers(s.cfg, s.apiLogger, kernelUC)

	app.Use(server_logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	kernelGroup := app.Group("kernel")
	mw := apiMiddlewares.NewMDWManager(s.cfg)

	kernelHttp.MapKernelRoutes(kernelGroup, kernelHandlers, mw)

	return nil
}
