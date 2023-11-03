package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// HTTPTest struct used to aux http table-test-driven
type HTTPTest struct {
	Description  string
	Route        string
	ExpectedCode int
	ExpectedBody interface{}
	BodyIsSlice  bool
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
	if httpTest.ExpectedBody == nil {
		return
	}

	var response interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Errorf("Err decoding JSON: %s", err)
	}

	assert.NotNil(t, response)

	var expected = structToMap(httpTest.ExpectedBody, httpTest.BodyIsSlice, t)
	var actual = structToMap(response, httpTest.BodyIsSlice, t)

	//Is actual response equals the expected
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected result:\nExpected: %s\nGot: %s", expected, actual)
	}
}

func structToMap(target interface{}, isSlice bool, t *testing.T) interface{} {
	var mapTarget interface{}

	jsonTarget, err := json.Marshal(target)
	if err != nil {
		t.Errorf("Erro ao decodificar o JSON esperado: %s", err)
	}

	if isSlice {
		var sliceMap []map[string]interface{}
		err = json.Unmarshal(jsonTarget, &sliceMap)
		mapTarget = sliceMap
	} else {
		var singleMap map[string]interface{}
		err = json.Unmarshal(jsonTarget, &singleMap)
		mapTarget = singleMap
	}

	if err != nil {
		t.Errorf("Erro ao decodificar o JSON esperado: %s", err)
	}

	return mapTarget
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
