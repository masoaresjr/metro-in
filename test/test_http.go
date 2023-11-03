package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

// HTTPTest struct used to aux http table-test-driven
type HTTPTest struct {
	Description  string
	Route        string
	ExpectedCode int
	ExpectedBody string
}

// RunHTTPTableDrivenTests runs all http tests
func RunHTTPTableDrivenTests(app *fiber.App, tests []HTTPTest, t *testing.T) {
	for _, httpTest := range tests {
		resp, _ := app.Test(httptest.NewRequest("GET", httpTest.Route, nil), -1)
		assert.Equalf(t, httpTest.ExpectedCode, resp.StatusCode, httpTest.Description)
	}
}
