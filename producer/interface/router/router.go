package router

import "github.com/gofiber/fiber/v2"

type WriteLogRoutes struct {
}

func NewWriteLogRoutes() *WriteLogRoutes {
	return &WriteLogRoutes{}
}

func (w *WriteLogRoutes) RegisterFavorite(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"foo": "bar"})
}
