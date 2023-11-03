package handler

import (
	"github.com/gofiber/fiber/v2"
	"metro-in/internal/common/customerrors"
)

type baseControllerImpl struct{}

// BaseController interface that deals with the http responses
type BaseController interface {
	RespondError(ctx *fiber.Ctx, err error) error
	RespondSuccessWithBody(ctx *fiber.Ctx, result interface{}) error
	RespondSuccessNoContent(ctx *fiber.Ctx) error
}

// NewBaseController constructor for StationController
func NewBaseController() BaseController {
	return &baseControllerImpl{}
}

func (h *baseControllerImpl) RespondError(c *fiber.Ctx, err error) error {
	if h.isHTTPError(err) {
		return c.Status(err.(customerrors.HTTPError).GetStatusCode()).SendString(err.Error())
	}

	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
}

func (h *baseControllerImpl) RespondSuccessWithBody(c *fiber.Ctx, result interface{}) error {
	return c.JSON(result)
}

func (h *baseControllerImpl) RespondSuccessNoContent(c *fiber.Ctx) error {
	c.Status(fiber.StatusNoContent)
	return nil
}

func (h *baseControllerImpl) isHTTPError(err error) bool {
	_, ok := err.(customerrors.HTTPError)
	return ok
}
