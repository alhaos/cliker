package webService

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type WebServerConfig struct {
	Address string `yaml:"address"`
	Mode    string `yaml:"mode"`
}

type (
	Service interface {
		Run()
	}

	service struct {
		conf   *WebServerConfig
		Router *gin.Engine
	}
)

func NewService(conf *WebServerConfig) (Service, error) {

	switch conf.Mode {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		return nil, fmt.Errorf("invalid webService mode value: %s", conf.Mode)
	}

	r := gin.Default()

	s := service{
		Router: r,
		conf:   conf,
	}

	s.registerRoutes(s.Router)
	err := s.loadTemplates(s.Router)
	if err != nil {
		panic(err)
	}

	return &s, nil
}
