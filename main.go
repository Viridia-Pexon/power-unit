package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"strings"
	"time"
	"viridia-power-unit/server"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	// disable tls check, only works for token refresh, not for proxied requests
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//load env

	/*
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("error godotenv: %s\n", err)
		}*/
}

func main() {

	app := fiber.New()
	server := &server.Powerunit{
		APP: app,
	}

	app.Use(recover.New()) // prevents crashes / reloads app on crash
	app.Use(logger.New())  // logs info of every request to console

	app.Use(cache.New(cache.Config{ // adds cacheing to requests (will relieve ArcGIS)
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   2 * time.Minute,
		CacheControl: true,
	}))

	allowOrigin := os.Getenv("ALLOW_ORIGIN_URL")
	if len(strings.TrimSpace(allowOrigin)) == 0 {
		allowOrigin = "http://localhost:5173" //default "local testing" cors (to prevents "wildcard cors error" )
	}

	// will compress the response depending on the Accept-Encoding header
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Use(cors.New(cors.Config{ //cors config
		Next:         nil,
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodOptions,
			fiber.MethodDelete,
			fiber.MethodPut,
			fiber.MethodPatch,
		}, ","),
		AllowHeaders:     "",
		AllowCredentials: true,
		ExposeHeaders:    "",
		MaxAge:           0,
	}))

	server.APP.Server().ReadBufferSize = 6400

	server.InitHealthEndpoint("/health")

	server.InitMetrics("/metrics")

	app.Listen(":8080")

	//move main to server.go in server package

}
