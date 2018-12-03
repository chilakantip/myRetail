package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chilakantip/avitar/pidfile"

	"github.com/chilakantip/avitar/log"
	"github.com/chilakantip/my_retail/env"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	e := echo.New()
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	assignHandlers(e)

	fmt.Println("Starting echo.labstack.com server...")
	go func() {
		if err := e.StartTLS(env.ServiceOnPort, "cert/cert.pem", "cert/key.pem"); err != nil {
			log.Info("got error,shutting down the server", err)
		}
	}()

	fmt.Println("Ready to serve the requests on the port", env.ServiceOnPort)
	fmt.Println("Setup OK.\nRunning... ")
	log.Info("Ready to serve the requests on the port", env.ServiceOnPort)
	log.Info("Setup OK.\nRunning... ")

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Info(err)
	}

	fmt.Println("shutting down the server... Done")

	pidfile.Drop()

	fmt.Println("Stopping goroutines... Done")
	fmt.Println("OK")
	log.Info("Shutdown OK")
	fmt.Println("All done. Bye Bye...")
	os.Exit(0)
}

func init() {
	doCommonSetUp()
}
