/*
Package views

Returns data and has internal checks for Package controllers
*/
package views

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ReturnMsg struct {
	Msg        string      `json:"msg"`
	Err        string      `json:"err,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
	Body       interface{} `json:"body,omitempty"`
}

func InvalidJson(c *fiber.Ctx, err error) error {
	status := fiber.StatusUnprocessableEntity
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "invalid json",
		Err:        fmt.Sprintf("%v", err),
		StatusCode: status,
	})
}

func NotFound(c *fiber.Ctx, err error) error {
	status := fiber.StatusNotFound
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "not found",
		Err:        fmt.Sprintf("%v", err),
		StatusCode: status,
	})
}

func AlreadyExists(c *fiber.Ctx, err error) error {
	status := fiber.StatusConflict
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "already exists",
		Err:        fmt.Sprintf("%v", err),
		StatusCode: status,
	})
}

func InternalServerError(c *fiber.Ctx, err error) error {

	status := fiber.StatusInternalServerError
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "something went wrong",
		Err:        "something went wrong",
		StatusCode: status,
	})
}

func Created(c *fiber.Ctx, body interface{}) error {
	status := fiber.StatusCreated
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "created",
		StatusCode: status,
		Body:       body,
	})
}

func OK(c *fiber.Ctx, body interface{}) error {
	status := fiber.StatusOK
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "ok",
		StatusCode: status,
		Body:       body,
	})
}

func Unauthorized(c *fiber.Ctx, err error) error {
	status := fiber.StatusUnauthorized
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "incorrect or missing credentials",
		Err:        fmt.Sprintf("%v", err),
		StatusCode: status,
	})
}

func Forbidden(c *fiber.Ctx, err error) error {
	status := fiber.StatusForbidden
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "you are forbidden from accessing this resource",
		Err:        fmt.Sprintf("%v", err),
		StatusCode: status,
	})
}

func InvalidQuery(c *fiber.Ctx, err error) error {
	status := fiber.StatusExpectationFailed
	return c.Status(status).JSON(ReturnMsg{
		Msg:        "invalid json",
		Err:        fmt.Sprintf("%v", err),
		StatusCode: status,
	})
}
