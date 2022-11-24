package hydra_api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
)

// AddProcessTimeHeader is a middleware which adds the time in nanoseconds it took to generate the response.
func AddProcessTimeHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		now := time.Now()
		defer func(ctx *fiber.Ctx) {
			latency := time.Since(now) / time.Millisecond
			ctx.Set("X-Process-Time", fmt.Sprintf("%dms", latency))
		}(c)
		return c.Next()
	}
}

// AddServerHeader is a middleware which adds a customisable `Server` header.
func AddServerHeader() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Server", "Hydra API")
		return c.Next()
	}
}

// middleware attaches HTTP middleware to the given Fiber application instance.
func (s *server) middleware() {
	s.app.Use(favicon.New())
	s.app.Use(etag.New(etag.Config{
		Weak: true,
	}))
	s.app.Use(logger.New())
	s.app.Use(helmet.New())
	s.app.Use(cors.New(cors.Config{
		AllowMethods:  "GET,OPTIONS,PUT,PATCH,POST,DELETE",
		ExposeHeaders: "Content-Type,Accept",
	}))
	s.app.Use(requestid.New())
	s.app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	s.app.Use(AddProcessTimeHeader())
	s.app.Use(AddServerHeader())
}
