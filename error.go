package ferror

import (
	"github.com/gofiber/fiber/v2"
)

type Error struct {
	FiberError *fiber.Error
	Cause      error
	Message    string
}
