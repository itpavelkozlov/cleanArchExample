package main

import (
	"cleanArch/pkg/logger"
)

func main()  {
	log := logger.NewLogger()
	log.Info("Application started")
}