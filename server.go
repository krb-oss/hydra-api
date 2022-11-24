package hydra_api

import (
	"fmt"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

type server struct{ app *fiber.App }

// Listen launches the server, attaching the HTTP handlers to their respective routes and serves requests from the
// configured port.
func (s *server) Listen() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// docker to bind to all available network interfaces
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("🛸 Listening on port %s\n", port)
	return s.app.Listen(addr)
}

// New is a factory which returns a new server object.
func New(app *fiber.App) *server {
	srv := &server{app}
	srv.middleware()
	srv.routes()
	return srv
}

// Fiber is a factory which returns a new Fiber application object.
func Fiber() *fiber.App {
	return fiber.New(fiber.Config{
		ErrorHandler:          ErrHandler,
		DisableStartupMessage: true,
		JSONEncoder:           sonic.Marshal,
		JSONDecoder:           sonic.Unmarshal,
	})
}

// ErrHandler returns a JSON formatted error responses.
func ErrHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"message": err.Error(),
	})
}
