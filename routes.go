package hydra_api

import "github.com/gofiber/fiber/v2"

func (s *server) routes() {
	// api := s.app.Group("/v1")
	s.app.Use(func(c *fiber.Ctx) error {
		return fiber.ErrNotFound
	})
}
