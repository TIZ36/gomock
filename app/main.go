package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"gomock/app/base"
	"gomock/app/internal/http"
	"gomock/app/internal/http/middleware"
	"gomock/types"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = types.DefaultConfigFilePath

func main() {

	// initialize gin router
	r := gin.New()
	r.Use(middleware.DefaultStructedLogger())
	r.Use(gin.Recovery())

	// load config
	cfg := loadConfig(configFile)
	log.Info().Msg(fmt.Sprintf("config is %v", cfg))

	// initialize app context
	base.NewAppContext(cfg)
	defer base.AppCtx.DestroyAppCtx()

	// define router
	r.GET("/chat", http.GetRouter)
	r.POST("/newKingdom", http.NewKingdom)

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

func loadConfig(path string) base.Config {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Error().Msg("read config error")
		panic(err)
	}

	var cfg base.Config

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		log.Error().Msg("invalid yaml format")
		panic(err)
	}

	return cfg
}
