package helpers

import "github.com/gofiber/fiber/v2"

type responseMessage struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"sucess"`
}

func SendMessage(c *fiber.Ctx, code int, message string, ok bool, data interface{}) error {

	response := responseMessage{
		Success: ok,
		Message: message,
		Data:    data,
	}

	return c.Status(code).JSON(response)
}
