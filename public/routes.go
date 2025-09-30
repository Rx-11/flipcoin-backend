package public

import (
	"github.com/Rx-11/prediction-backend/common"
	"github.com/gofiber/fiber/v2"
)

func fetchCurrentPrice(c *fiber.Ctx) error {
	price, err := fetchPrice()
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).SendString("Failed to get price: " + err.Error())
	}
	return c.JSON(fiber.Map{
		"price": price.String(),
	})
}

func fetchRoundDetails(c *fiber.Ctx) error {
	var body fetchDetailsStruct
	err := c.QueryParser(&body)
	if err != nil {
		return c.Status(common.ErrInvalidParams.StatusCode).SendString("Failed to parse request body: " + err.Error())
	}
	round, err := fetchRound(body.Epoch)
	if err != nil {
		return c.Status(common.ErrInternalServerError.StatusCode).SendString("Failed to get round: " + err.Error())
	}
	return c.JSON(fiber.Map{
		"epoch":          round.Epoch.String(),
		"startTimestamp": round.StartTimestamp.String(),
		"lockTimestamp":  round.LockTimestamp.String(),
		"closeTimestamp": round.CloseTimestamp.String(),
		"lockPrice":      round.LockPrice.String(),
		"closePrice":     round.ClosePrice.String(),
		"totalAmount":    round.TotalAmount.String(),
		"bullAmount":     round.BullAmount.String(),
		"bearAmount":     round.BearAmount.String(),
	})
}
