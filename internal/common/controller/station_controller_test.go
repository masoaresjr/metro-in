package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"metro-in/internal/common/service/mock"
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

func TestGetByName(t *testing.T) {

	stationController := stationControllerImpl{ stationService: mock.NewStationServiceMock() }

	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			route:        "/station/corinthians-itaquera?isLogged=true",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route is incomplete",
			route:        "/station/",
			expectedCode: 404,
		},
		{
			description:  "get HTTP status 401, when user is not logged",
			route:        "/station/corinthians-itaquera?isLogged=false",
			expectedCode: 401,
		},
	}

	app := fiber.New()
	app.Get("/station/:name", authMiddleware, stationController.GetByName)

	for _, test := range tests {
		resp, _ := app.Test(httptest.NewRequest("GET", test.route, nil), -1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func authMiddleware(c *fiber.Ctx) error {

	//isLogged is at the queryParams just for tests purposes
	//at the real scenario there are sessions and tokens
	isUserLoggedIn := c.Query("isLogged")

	if isUserLoggedIn == "false" {
		return c.Status(fiber.StatusUnauthorized).SendString("Acesso não autorizado")
	}

	return c.Next()
}