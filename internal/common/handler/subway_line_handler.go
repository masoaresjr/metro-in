package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"metro-in/internal/common/customerrors"
	"metro-in/internal/common/service"
	"strconv"
)

// subwayLineControllerImpl implementation for SubwayLineController
type subwayLineControllerImpl struct {
	baseController BaseController
	subwayService  service.SubwayLineService
}

// SubwayLineController interface for subway lines routes
type SubwayLineController interface {
	GetAll(*fiber.Ctx) error
	GetByID(*fiber.Ctx) error
	GetByCompanyID(*fiber.Ctx) error
}

// NewSubwayLineController constructor for SubwayLineController
func NewSubwayLineController(dbClient *gorm.DB) SubwayLineController {
	return &subwayLineControllerImpl{subwayService: service.NewSubwayLineService(dbClient), baseController: &baseControllerImpl{}}
}

// GetAll godoc
func (h *subwayLineControllerImpl) GetAll(c *fiber.Ctx) error {
	response, err := h.subwayService.GetAll()
	if err != nil {
		return h.baseController.RespondError(c, err)
	}
	return h.baseController.RespondSuccessWithBody(c, response)
}

// GetByID godoc
func (h *subwayLineControllerImpl) GetByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return h.baseController.RespondError(
			c, &customerrors.InvalidParameterError{ParameterType: "integer", ParameterName: "id"},
		)
	}

	response, err := h.subwayService.GetByID(uint(id))
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
			c, &customerrors.InvalidParameterError{ParameterType: "integer", ParameterName: "company_id"},
		)
	}

	response, err := h.subwayService.GetByID(uint(id))
	if err != nil {
		return h.baseController.RespondError(c, err)
	}

	return h.baseController.RespondSuccessWithBody(c, response)
}
