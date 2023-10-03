package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Arbitration    ArbitrationConfig
	BinanceAccount BinanceAccount
	Logger         Logger
	RabbitMQ       RabbitMQ
	System         SystemConfig
	Server         ServerConfig
	GarantexAPI    GarantexAPI
	TerminalAccess TerminalAccess
}

type GarantexAPI struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
	UID        string `json:"UID"`
}

type TerminalAccess struct {
	APIKey string
}

type RabbitMQ struct {
	Access struct {
		User     string
		Password string
		Host     string
		Port     string
	}
	Queues struct {
		OrderStatusUpdated string
	}
}

type SystemConfig struct {
	MaxGoRoutines int64
}

type ServerConfig struct {
	AppVersion                  string `json:"appVersion"`
	Host                        string `json:"host" validate:"required"`
	Port                        string `json:"port" validate:"required"`
	ShowUnknownErrorsInResponse bool   `json:"showUnknownErrorsInResponse"`
}

type ArbitrationConfig struct {
	MaxDepth      int64
	MinimalMargin float64
	MinInvesting  float64
	MaxInvesting  float64
}

type Logger struct {
	Level          string   `json:"level"`
	SkipFrameCount int      `json:"skipFrameCount"`
	InFile         bool     `json:"inFile"`
	FilePath       string   `json:"filePath"`
	InTG           bool     `json:"inTg"`
	ChatID         int64    `json:"chatID"`
	TGToken        string   `json:"-"`
	AlertUsers     []string `json:"alertUsers"`
}

type BinanceAccount struct {
	APIKey    string
	APISecret string
}

func LoadConfig() (*viper.Viper, error) {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
