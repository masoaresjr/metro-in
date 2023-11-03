package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"metro-in/internal/common/service/mock"
	"metro-in/test"
	"net/http/httptest"
	"testing"
)

// TestNewStationController test if SubwayLineController constructor is returning the correct implementation
func TestNewStationController(t *testing.T) {
	stationController := NewStationController(nil)

	// stationController can't be null
	assert.NotNil(t, stationController)

	// stationController must implement SubwayLineController interface
	assert.Implements(t, (*StationController)(nil), stationController)
}

func TestStationControllerGetByName(t *testing.T) {

	stationController := stationControllerImpl{ stationService: mock.NewStationServiceMock() }

	tests := []test.HttpTest {
		{
			Description:  "get HTTP status 200",
			Route:        "/station/corinthians-itaquera?isLogged=true",
			ExpectedCode: 200,
		},
		{
			Description:  "get HTTP status 404, when route is incomplete",
			Route:        "/station/",
			ExpectedCode: 404,
		},
		{
			Description:  "get HTTP status 401, when user is not logged",
			Route:        "/station/corinthians-itaquera?isLogged=false",
			ExpectedCode: 401,
		},
		{
			Description:  "Test service error",
			Route:        "/station/patriarca?isLogged=true",
			ExpectedCode: 500,
		},
	}

	app := fiber.New()
	app.Get("/station/:name", authMiddleware, stationController.GetStationByName)

	for _, httpTest := range tests {
		resp, _ := app.Test(httptest.NewRequest("GET", httpTest.Route, nil), -1)
		assert.Equalf(t, httpTest.ExpectedCode, resp.StatusCode, httpTest.Description)
	}
}

func authMiddleware(c *fiber.Ctx) error {

	//IsLogged is inside the queryParams just for tests purposes.
	//At the real scenario there are sessions and tokens
	isUserLoggedIn := c.Query("isLogged")

	if isUserLoggedIn == "false" {
		return c.Status(fiber.StatusUnauthorized).SendString("Acesso não autorizado")
	}

	return c.Next()
}