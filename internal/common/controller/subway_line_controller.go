package controller

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
func (c *subwayLineControllerImpl) GetAll(ctx *fiber.Ctx) error {
	response, err := c.subwayService.GetAll()
	if err != nil {
		return c.baseController.RespondError(ctx, err)
	}
	return c.baseController.RespondSuccessWithBody(ctx, response)
}

// GetByID godoc
func (c *subwayLineControllerImpl) GetByID(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return c.baseController.RespondError(
			ctx, &customerrors.InvalidParameterError{ ParameterType: "integer", ParameterName: "id" },
		)
	}

	response, err := c.subwayService.GetByID(uint(id))
	if err != nil {
		return c.baseController.RespondError(ctx, err)
	}

	return c.baseController.RespondSuccessWithBody(ctx, response)
}

// GetByCompanyID godoc
func (c *subwayLineControllerImpl) GetByCompanyID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("company_id"))
	if err != nil {
		return c.baseController.RespondError(
			ctx, &customerrors.InvalidParameterError{ ParameterType: "integer", ParameterName: "company_id" },
		)
	}

	response, err := c.subwayService.GetByID(uint(id))
	if err != nil {
		return c.baseController.RespondError(ctx, err)
	}

	return c.baseController.RespondSuccessWithBody(ctx, response)
}
