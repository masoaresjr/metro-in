package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HTTPTest struct used to aux http table-test-driven
type HTTPTest struct {
	Description  string
	Route        string
	ExpectedCode int
	ExpectedBody string
	ExpectedInterface interface{}
}

// RunHTTPTableDrivenTests runs all http tests
func RunHTTPTableDrivenTests(app *fiber.App, tests []HTTPTest, t *testing.T) {
	for _, httpTest := range tests {
		resp, _ := app.Test(httptest.NewRequest("GET", httpTest.Route, nil), -1)

		validateBody(resp, httpTest, t)

		assert.Equalf(t, httpTest.ExpectedCode, resp.StatusCode, httpTest.Description)
	}
}

func validateBody(resp *http.Response, httpTest HTTPTest, t *testing.T) {
	if httpTest.ExpectedInterface == nil {
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&httpTest.ExpectedInterface); err != nil {
		t.Errorf("Erro ao decodificar o JSON: %s", err)
	}

	assert.NotNil(t, httpTest.ExpectedInterface)
}

// AuthMiddleware used to test result depending on user current login status
func AuthMiddleware(c *fiber.Ctx) error {

	//IsLogged is inside the queryParams just for tests purposes.
	//At the real scenario there are sessions and tokens
	isUserLoggedIn := c.Query("isLogged")

	if isUserLoggedIn != "true" {
		return c.Status(fiber.StatusUnauthorized).SendString("Acesso não autorizado")
	}

	return c.Next()
}
