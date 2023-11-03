package controller

import (
	"github.com/gofiber/fiber/v2"
	"metro-in/internal/common/customerrors"
)

type baseControllerImpl struct{}

type BaseController interface {
	RespondError(ctx *fiber.Ctx, err error) error
	RespondSuccessWithBody(ctx *fiber.Ctx, result interface{}) error
	RespondSuccessNoContent(ctx *fiber.Ctx) error
}

// NewBaseController constructor for StationController
func NewBaseController() BaseController {
	return &baseControllerImpl{}
}

func (c *baseControllerImpl) RespondError(ctx *fiber.Ctx, err error) error {
	if c.isHttpError(err) {
		return ctx.Status(err.(customerrors.HttpError).GetStatusCode()).SendString(err.Error())
	}

	return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
}

func (c *baseControllerImpl) RespondSuccessWithBody(ctx *fiber.Ctx, result interface{}) error {
	return ctx.JSON(result)
}

func (c *baseControllerImpl) RespondSuccessNoContent(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNoContent)
	return nil
}

func (c *baseControllerImpl) isHttpError(err error) bool {
	_, ok := err.(customerrors.HttpError)
	return ok
}
