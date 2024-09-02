package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"metro-in/internal/common/constants"
	"metro-in/internal/common/customerrors"
	"metro-in/internal/common/middleware"
	"metro-in/internal/common/service"
	"strconv"
)

// subwayLineControllerImpl implementation for SubwayLineController
type subwayLineControllerImpl struct {
	baseController    BaseController
	subwayLineService service.SubwayLineService
}

// SubwayLineController interface for subway lines routes
type SubwayLineController interface {
	RegisterRoutes(app *fiber.App)
	GetAll(*fiber.Ctx) error
	GetByID(*fiber.Ctx) error
	GetByCompanyID(*fiber.Ctx) error
}

// NewSubwayLineController constructor for SubwayLineController
func NewSubwayLineController(dbClient *gorm.DB) SubwayLineController {
	return &subwayLineControllerImpl{subwayLineService: service.NewSubwayLineService(dbClient), baseController: &baseControllerImpl{}}
}

// RegisterRoutes set the http routes for the handler
func (h *subwayLineControllerImpl) RegisterRoutes(app *fiber.App) {
	lineRouter := app.Group(constants.LineHandler)
	lineRouter.Get("/", middleware.AuthMiddleware, h.GetAll)
}

// GetAll godoc
func (h *subwayLineControllerImpl) GetAll(c *fiber.Ctx) error {
	response, err := h.subwayLineService.GetAll()
	if err != nil {
		return h.baseController.RespondError(c, err)
	}
	return h.baseController.RespondSuccessWithBody(c, response)
}

// GetByID godoc
func (h *subwayLineControllerImpl) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params(constants.ID))
	if err != nil {
		return h.baseController.RespondError(
			c, &customerrors.InvalidParameterError{ParameterType: constants.Integer, ParameterName: constants.ID},
		)
	}

	response, err := h.subwayLineService.GetByID(uint(id))
	if err != nil {
		return h.baseController.RespondError(c, err)
	}

	return h.baseController.RespondSuccessWithBody(c, response)
}

// GetByCompanyID godoc
func (h *subwayLineControllerImpl) GetByCompanyID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("company_id"))
	if err != nil {
		return h.baseController.RespondError(
			c, &customerrors.InvalidParameterError{ParameterType: constants.Integer, ParameterName: constants.CompanyID},
		)
	}

	response, err := h.subwayLineService.GetByCompanyID(uint(id))
	if err != nil {
		return h.baseController.RespondError(c, err)
	}

	return h.baseController.RespondSuccessWithBody(c, response)
}
