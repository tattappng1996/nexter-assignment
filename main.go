package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"time"

	publicHandler "nexter-assignment/handler/http/public"
	"nexter-assignment/repository"
	"nexter-assignment/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Nexter-Assignment Example API
// @version 1.0
// @description This is a nexter-assignment server.

// @contact.name tattapong.kun@gmail.com

var cashRegisterFilePath = "./save/cash_register.json"
var cashStorageFilePath = "./save/cash_storage.json"

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper: middleware.DefaultSkipper,
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAuthorization, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlRequestHeaders,
			"sec-ch-ua", "sec-ch-ua-mobile", "sec-ch-ua-platform", "User-Agent", "X-Session-Id"},
		AllowCredentials: true,
	}))
	// e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"time_now": time.Now(),
			"succuss":  true,
		})
	})

	uc := usecase.NewUsecase(
		repository.NewRepository(cashRegisterFilePath, cashStorageFilePath),
	)

	examGroup := e.Group("/exam")

	publicHandler.NewRouter(examGroup, uc)
	log.Println("start exam http public ... ")

	serveGracefulShutdown(e)
}

func serveGracefulShutdown(e *echo.Echo) {
	go func() {
		if err := e.Start(":8080"); err != nil {
			fmt.Println("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gracefulShutdownTimeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err.Error())
	}
}

func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
		panic(v)
	}
	fmt.Println(fmt.Sprintf("%s ==>  %s", k, v))
	return v
}
