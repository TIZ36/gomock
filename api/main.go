package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gomock/api/lib"
	"gomock/api/lib/config"
	"gomock/api/lib/ctx"
	"gomock/api/lib/middleware"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = "../config/dev.yaml"

func main() {

	// initialize gin router
	r := gin.New()
	r.Use(middleware.DefaultStructedLogger())
	r.Use(gin.Recovery())

	// load config
	cfg := loadConfig(configFile)
	log.Info().Msg(fmt.Sprintf("config is %v", cfg))

	// initialize app context
	ctx.NewAppContext(cfg)
	defer ctx.AppCtx.DestroyAppCtx()

	// define router
	r.GET("/chat", lib.GetRouter)

	// set up logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	gin.DefaultWriter = log.Logger

	err := r.Run(":8888")
	if err != nil {
		return
	}

	fmt.Println("hhh")
}

func loadConfig(path string) config.Config {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Error().Msg("read config error")
		panic(err)
	}

	var cfg config.Config

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		log.Error().Msg("invalid yaml format")
		panic(err)
	}

	return cfg
}
