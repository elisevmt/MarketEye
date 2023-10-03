package server

import (
	"MarketEye/config"
	"MarketEye/pkg/httpErrors"
	"MarketEye/pkg/logger"
	"fmt"
	//"github.com/go-playground/locales/tg"
	//"github.com/go-playground/locales/tg"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// Server struct
type Server struct {
	fiber     *fiber.App
	cfg       *config.Config
	database  *sqlx.DB
	apiLogger *logger.ApiLogger
}

func NewServer(cfg *config.Config, apiLogger *logger.ApiLogger) *Server {
	return &Server{
		fiber:     fiber.New(fiber.Config{ErrorHandler: httpErrors.Handler, DisableStartupMessage: true}),
		cfg:       cfg,
		apiLogger: apiLogger,
	}
}

func (s *Server) Run() error {
	if err := s.MapHandlers(s.fiber, s.apiLogger); err != nil {
		s.apiLogger.Fatalf("Cannot map handlers: ", err)
	}
	s.apiLogger.Infof("Start server on port: %s:%s", s.cfg.Server.Host, s.cfg.Server.Port)
	if err := s.fiber.Listen(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil {
		s.apiLogger.Fatalf("Error starting Server: ", err)
	}
	return nil
}
