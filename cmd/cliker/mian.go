package main

import (
	"flag"
	"github.com/alhaos/cliker/internal/config"
	"github.com/alhaos/cliker/internal/webService"
)

func main() {

	// init config
	configFilename := ParseFileName()

	conf, err := config.FromFile(configFilename)
	if err != nil {
		panic(err)
	}

	// init webService
	ws, err := webService.NewService(conf.WebServer)
	if err != nil {
		panic(err)
	}

	// run web service
	ws.Run()
}

// ParseFileName take config file name
func ParseFileName() string {
	filenamePointer := flag.String("config", "config/config.yml", "Cliker config file")
	flag.Parse()
	filename := *filenamePointer
	return filename
}
