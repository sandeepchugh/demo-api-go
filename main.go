package main

import (
	"github.com/sandeepchugh/banking/app"
	"github.com/sandeepchugh/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
