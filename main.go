package main

import (
	"flag"

	"github.com/lucasponce/jaeger-proto-client/config"
	"github.com/lucasponce/jaeger-proto-client/jaeger"
	"github.com/lucasponce/jaeger-proto-client/log"
)

// Command line arguments
var (
	argConfigFile = flag.String("config", "", "Path to the YAML configuration file. If not specified, environment variables will be used for configuration.")
)

func main() {
	log.InitializeLogger()
	flag.Parse()

	// load config file if specified, otherwise, rely on environment variables to configure us
	if *argConfigFile != "" {
		c, err := config.LoadFromFile(*argConfigFile)
		if err != nil {
			log.Fatal(err)
		}
		config.Set(c)
	} else {
		log.Infof("No configuration file specified. Will rely on environment for configuration.")
		config.Set(config.NewConfig())
	}
	log.Infof("Jaeger Proto Client")

	token := ""
	jaegerClient, err := jaeger.NewClient(token)
	if err != nil {
		log.Errorf("NewClient error: %s", err)
		return
	}
	status, err2 := jaegerClient.GetServiceStatus()
	if err2 != nil {
		log.Errorf("GetServiceStatus error: %s", err2)
	}
	log.Infof("GetServiceStatus %t", status)
}
