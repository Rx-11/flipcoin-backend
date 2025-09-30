package public

import (
	"github.com/Rx-11/flipcoin-backend/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func MountRoutes(router *fiber.App, wsBroadcaster *ws.WsBroadcaster) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	apiGroup := router.Group("/api")

	apiGroup.Get("/price", fetchCurrentPrice)
	apiGroup.Get("/round", fetchRoundDetails)

	router.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Get("/ws/updates", websocket.New(func(conn *websocket.Conn) {
		wsBroadcaster.RegisterClient(conn)
		wsBroadcaster.ClientLoop(conn)
	}))

}
