package controllers

import (
	"net/http"
	"white-goods-multifinace/dto"
	"white-goods-multifinace/models"
	"white-goods-multifinace/repositories"
	"white-goods-multifinace/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ItemController struct {
	itemRepo  repositories.ItemRepository
	itemTenor repositories.ItemTenorRepository
}

func NewItemController(itemRepo repositories.ItemRepository, itemTenor repositories.ItemTenorRepository) *ItemController {
	return &ItemController{
		itemRepo:  itemRepo,
		itemTenor: itemTenor,
	}
}

func (ic *ItemController) CreateItem(c echo.Context) error {
	var itemBody dto.AddItemBody

	if err := c.Bind(&itemBody); err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid request body"))
	}

	if itemBody.Name == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Name is required"))
	}

	if itemBody.NormalPrice == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Normal price is required"))
	}

	if itemBody.AdminFee == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("Admin fee is required"))
	}

	if itemBody.OTRID == 0 {
		return utils.HandlerError(c, utils.NewBadRequestError("OTR ID is required"))
	}

	if itemBody.Tenor1ID == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Tenor 1 ID is required"))
	}

	if itemBody.Tenor2ID == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Tenor 2 ID is required"))
	}

	if itemBody.Tenor3ID == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Tenor 3 ID is required"))
	}

	if itemBody.Tenor6ID == "" {
		return utils.HandlerError(c, utils.NewBadRequestError("Tenor 6 ID is required"))
	}

	parsedTenor1ID, err := uuid.Parse(itemBody.Tenor1ID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid tenor 1 ID"))
	}

	parsedTenor2ID, err := uuid.Parse(itemBody.Tenor2ID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid tenor 2 ID"))
	}

	parsedTenor3ID, err := uuid.Parse(itemBody.Tenor3ID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid tenor 3 ID"))
	}

	parsedTenor6ID, err := uuid.Parse(itemBody.Tenor6ID)
	if err != nil {
		return utils.HandlerError(c, utils.NewBadRequestError("Invalid tenor 6 ID"))
	}

	item := models.Item{
		Name:        itemBody.Name,
		NormalPrice: itemBody.NormalPrice,
		AdminFee:    itemBody.AdminFee,
		OTRID:       itemBody.OTRID,
	}

	itemID, err := ic.itemRepo.CreateItem(&item)
	if err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to create item"))
	}

	itemTenor := []models.ItemTenor{
		{
			ItemID:   itemID,
			TenorID:  parsedTenor1ID,
			Interest: itemBody.InterestTenor1,
		},
		{
			ItemID:   itemID,
			TenorID:  parsedTenor2ID,
			Interest: itemBody.InterestTenor2,
		},
		{
			ItemID:   itemID,
			TenorID:  parsedTenor3ID,
			Interest: itemBody.InterestTenor3,
		},
		{
			ItemID:   itemID,
			TenorID:  parsedTenor6ID,
			Interest: itemBody.InterestTenor6,
		},
	}

	if err := ic.itemTenor.CreateItemLimit(&itemTenor); err != nil {
		return utils.HandlerError(c, utils.NewInternalError("Failed to create item tenor"))
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Item created successfully",
	})

}
