package handler

import (
	"errors"
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
	var statusCode = fiber.StatusInternalServerError

	if h.isHTTPError(err) {
		statusCode = err.(customerrors.HTTPError).GetStatusCode()
	}

	if statusCode == fiber.StatusInternalServerError {
		//persist log
	}

	return c.Status(statusCode).SendString(err.Error())
}

func (h *baseControllerImpl) RespondSuccessWithBody(c *fiber.Ctx, result interface{}) error {
	return c.JSON(result)
}

func (h *baseControllerImpl) RespondSuccessNoContent(c *fiber.Ctx) error {
	c.Status(fiber.StatusNoContent)
	return nil
}

func (h *baseControllerImpl) isHTTPError(err error) bool {
	var HTTPError customerrors.HTTPError
	return errors.As(err, &HTTPError)
}
