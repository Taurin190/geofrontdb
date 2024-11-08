package main

import (
	"flag"
	"fmt"

	"github.com/Taurin190/geofrontdb/internal/config"
	log "github.com/sirupsen/logrus"
)

func init() {
	flag.Parse()
}

func main() {
	log.Info("Starting Geofront Server...")
	config := config.NewServerConfig()
	log.Infof("Server started on port %s", config.Port)
	fmt.Println("Hello, World!")
}
