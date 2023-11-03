package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"metro-in/internal/common/customerrors"
	"metro-in/test"
	"testing"
)

// TestBaseController test if BaseController constructor is returning the correct implementation
func TestBaseController(t *testing.T) {
	baseController := NewBaseController()

	// baseController can't be null
	assert.NotNil(t, baseController)

	// stationController must implement SubwayLineController interface
	assert.Implements(t, (*BaseController)(nil), baseController)
}

func TestBaseControllerRespondError(t *testing.T) {

	stationController := baseControllerImpl{}

	tests := []controller.HTTPTest{
		{
			Description:  "get HTTP status 404, when route is incomplete",
			Route:        "/fake-route/http-error",
			ExpectedCode: 400,
		},
		{
			Description:  "get HTTP status 401, when user is not logged",
			Route:        "/fake-route/unknown-error",
			ExpectedCode: 500,
		},
	}

	app := fiber.New()

	app.Get("/fake-route/http-error", func(c *fiber.Ctx) error {
		return stationController.RespondError(c, &customerrors.EmptyParameterError{ ParameterName: "id" })
	})

	app.Get("/fake-route/unknown-error", func(c *fiber.Ctx) error {
		return stationController.RespondError(c, errors.New("unexpected error"))
	})

	controller.RunHttpTableDrivenTests(app, tests, t)
}

func TestBaseControllerRespondSuccess(t *testing.T) {

	stationController := baseControllerImpl{}

	tests := []controller.HTTPTest{
		{
			Description:  "get HTTP status 200",
			Route:        "/fake-route",
			ExpectedCode: 200,
		},
	}

	app := fiber.New()

	app.Get("/fake-route", func(c *fiber.Ctx) error {
		return stationController.RespondSuccessWithBody(c, map[string]string{ "Test" : "Test" })
	})

	controller.RunHttpTableDrivenTests(app, tests, t)
}
