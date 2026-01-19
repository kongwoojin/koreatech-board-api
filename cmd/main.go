package main

import (
	"flag"
	"fmt"
	_ "koreatech-board-api/cmd/docs"
	"koreatech-board-api/cmd/routes"
	"koreatech-board-api/cmd/utils"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			KOREATECH board REST API
//	@version		1.0
//	@description	This is unofficial version of KOREATECH board REST API

//	@contact.name	Developer
//	@contact.email	kongwoojin03@gmail.com

// @BasePath	/v3
func main() {
	// Parse flags
	host := flag.String("h", "0.0.0.0", "Set host")
	port := flag.String("p", "1323", "Set port")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	// Echo instance
	e := echo.New()

	if utils.IsRunningInContainer() {
		// If running in container, log to stdout
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339}: ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, user_agent=${user_agent}\n",
			Output: os.Stdout,
		}))
	} else {
		// If running in local, log to file
		f, err := os.Create("access.log")

		if err != nil {
			panic(err)
		}

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${time_rfc3339}: ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, user_agent=${user_agent}\n",
			Output: f,
		}))
	}

	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.APIRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprint(*host, ":", *port)))
}
