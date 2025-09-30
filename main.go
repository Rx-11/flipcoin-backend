package main

import (
	"log"

	"github.com/Rx-11/flipcoin-backend/config"
	"github.com/Rx-11/flipcoin-backend/public"
	"github.com/Rx-11/flipcoin-backend/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "flipcoin-backend",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	config.Init()
	log.Println("Loaded configs.")
	config.InitContract()
	log.Println("Loaded contract.")

	wsBroadcaster := ws.NewWsBroadcaster()
	go wsBroadcaster.Run()

	public.MountRoutes(app, wsBroadcaster)

	go public.RoundManager(wsBroadcaster)
	log.Println("Server started at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))

}
