package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"metro-in/internal/common/entity"
	"metro-in/internal/common/service/mock"
	controller "metro-in/test"
	"testing"
)

// TestNewSubwayLineController test if SubwayLineController constructor is returning the correct implementation
func TestNewSubwayLineController(t *testing.T) {
	subwayLineController := NewSubwayLineController(nil)

	// subwayLineController can't be null
	assert.NotNil(t, subwayLineController)

	// subwayLineController must implement SubwayLineController interface
	assert.Implements(t, (*SubwayLineController)(nil), subwayLineController)
}

func TestSubwayLineControllerImpl_GetAll(t *testing.T) {
	subwayLineController := subwayLineControllerImpl{subwayLineService: mock.NewSubwayLineServiceMock(), baseController: NewBaseController()}
	tests := []controller.HTTPTest{
		{
			Description:  "get HTTP status 200",
			Route:        "/line?isLogged=true",
			ExpectedCode: 200,
			BodyIsSlice:  true,
			ExpectedBody: []entity.SubwayLine{
				{
					Model: gorm.Model{
						ID: 1,
					},
					CompanyID: 2,
					Name:      "Vermelha",
				},
			},
		},
		{
			Description:  "get HTTP status 401 when user is not logged",
			Route:        "/line",
			ExpectedCode: 401,
		},
	}

	app := fiber.New()
	app.Get("/line", controller.AuthMiddleware, subwayLineController.GetAll)

	controller.RunHTTPTableDrivenTests(app, tests, t)

	//Error scenario
	subwayLineController = subwayLineControllerImpl{subwayLineService: mock.NewSubwayLineServiceMockError(), baseController: NewBaseController()}
	tests = []controller.HTTPTest{
		{
			Description:  "get HTTP status 200",
			Route:        "/line?isLogged=true",
			ExpectedCode: 500,
		},
	}
	controller.RunHTTPTableDrivenTests(app, tests, t)
}

func TestSubwayLineControllerImpl_GetByID(t *testing.T) {
	subwayLineController := subwayLineControllerImpl{subwayLineService: mock.NewSubwayLineServiceMock(), baseController: NewBaseController()}
	tests := []controller.HTTPTest{
		{
			Description:  "get HTTP status 200",
			Route:        "/line/1?isLogged=true",
			ExpectedCode: 200,
			BodyIsSlice:  false,
			ExpectedBody: entity.SubwayLine{
				Model: gorm.Model{
					ID: 1,
				},
				CompanyID: 2,
				Name:      "Vermelha",
			},
		}, {
			Description:  "get HTTP status 400 wrong parameter type",
			Route:        "/line/meu-id?isLogged=true",
			ExpectedCode: 400,
		},
		{
			Description:  "get HTTP status 401 when user is not logged",
			Route:        "/line/1",
			ExpectedCode: 401,
		},
	}

	app := fiber.New()
	app.Get("/line/:id", controller.AuthMiddleware, subwayLineController.GetByID)

	controller.RunHTTPTableDrivenTests(app, tests, t)

	//Error scenario
	subwayLineController = subwayLineControllerImpl{subwayLineService: mock.NewSubwayLineServiceMockError(), baseController: NewBaseController()}
	tests = []controller.HTTPTest{
		{
			Description:  "get HTTP status 200",
			Route:        "/line/1?isLogged=true",
			ExpectedCode: 500,
		},
	}
	controller.RunHTTPTableDrivenTests(app, tests, t)
}

func TestSubwayLineControllerImpl_GetByCompanyID(t *testing.T) {
	subwayLineController := subwayLineControllerImpl{subwayLineService: mock.NewSubwayLineServiceMock(), baseController: NewBaseController()}
	tests := []controller.HTTPTest{
		{
			Description:  "get HTTP status 200",
			Route:        "/line/company/1?isLogged=true",
			ExpectedCode: 200,
			BodyIsSlice:  true,
			ExpectedBody: []entity.SubwayLine{
				{
					Model: gorm.Model{
						ID: 1,
					},
					CompanyID: 2,
					Name:      "Vermelha",
				},
			},
		}, {
			Description:  "get HTTP status 400 wrong parameter type",
			Route:        "/line/company/meu-id?isLogged=true",
			ExpectedCode: 400,
		},
		{
			Description:  "get HTTP status 401 when user is not logged",
			Route:        "/line/company/1",
			ExpectedCode: 401,
		},
	}

	app := fiber.New()
	app.Get("/line/company/:company_id", controller.AuthMiddleware, subwayLineController.GetByCompanyID)

	controller.RunHTTPTableDrivenTests(app, tests, t)

	//Error scenario
	subwayLineController = subwayLineControllerImpl{subwayLineService: mock.NewSubwayLineServiceMockError(), baseController: NewBaseController()}
	tests = []controller.HTTPTest{
		{
			Description:  "get HTTP status 200",
			Route:        "/line/company/1?isLogged=true",
			ExpectedCode: 500,
		},
	}
	controller.RunHTTPTableDrivenTests(app, tests, t)
}
