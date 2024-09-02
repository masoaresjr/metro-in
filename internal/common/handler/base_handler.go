package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"metro-in/internal/common/constants"
	"metro-in/internal/common/customerrors"
	"strconv"
)

type baseControllerImpl struct{}

// BaseController interface that deals with the http responses
type BaseController interface {
	GetUintParam(c *fiber.Ctx, paramName string) (paramValue uint, err error)
	GetStringParam(c *fiber.Ctx, paramName string) (paramValue string, err error)

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

func (h *baseControllerImpl) GetUintParam(c *fiber.Ctx, paramName string) (paramValue uint, err error) {
	param := c.Params(paramName)
	if param == "" {
		return 0, h.RespondError(c, customerrors.NewEmptyParameterError(constants.ID))
	}

	uintParam, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, h.RespondError(c, customerrors.NewInvalidParameterError(constants.Integer, paramName))
	}

	return uint(uintParam), nil
}

func (h *baseControllerImpl) GetStringParam(c *fiber.Ctx, paramName string) (paramValue string, err error) {
	param := c.Params(paramName)
	if param == "" {
		return "", h.RespondError(c, customerrors.NewEmptyParameterError(paramName))
	}
	return param, nil
}

func (h *baseControllerImpl) isHTTPError(err error) bool {
	var HTTPError customerrors.HTTPError
	return errors.As(err, &HTTPError)
}
