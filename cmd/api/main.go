package main

import (
	"MarketEye/config"
	"MarketEye/internal/server"
	"MarketEye/pkg/httpErrors"
	"MarketEye/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"log"
)

func main() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Println("Starting server")
	v, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Cannot cload config: ", err.Error())
	}
	cfg, err := config.ParseConfig(v)
	if err != nil {
		log.Fatalf("Config parse error", err.Error())
	}
	log.Println("Config loaded")
	appLogger := logger.NewApiLogger(cfg)
	err = appLogger.InitLogger()
	if err != nil {
		log.Fatalf("Cannot init logger: %v", err.Error())
	}
	appLogger.Infof("Logger successfully started with - Level: %s, InFile: %t (filePath: %s), InTG: %t (chatID: %d)",
		cfg.Logger.Level,
		cfg.Logger.InFile,
		cfg.Logger.FilePath,
		cfg.Logger.InTG,
		cfg.Logger.ChatID)

	httpErrors.Init(cfg, appLogger)
	s := server.NewServer(cfg, appLogger)
	if err = s.Run(); err != nil {
		appLogger.ErrorFull(err)
	}
}
