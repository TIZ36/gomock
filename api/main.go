package main

import (
	"errors"
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
	testDataStructure()
	r := gin.New()

	r.Use(middleware.DefaultStructedLogger())
	r.Use(gin.Recovery())

	cfg := loadConfig(configFile)

	log.Info().Msg(fmt.Sprintf("config is %v", cfg))
	ctx.NewAppContext(cfg)

	defer ctx.AppCtx.DestroyAppCtx()

	r.GET("/chat", lib.GetRouter)

	log.Print("prepare to listen on 8888")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	gin.DefaultWriter = log.Logger

	err := errors.New("mock err here")
	log.Error().Stack().Err(err).Msg("a info log")
	err = r.Run(":8888")
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

func testDataStructure() {
	//myStack := datastructure.New[int32]()
	//
	//fmt.Println(myStack.Pop())
	//myStack.Push(44)
	//fmt.Println(myStack.Pop())

}
