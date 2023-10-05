package main

import (

	//"uts/database"
	"log"
	"uts/database"
	"uts/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	//Kumpulan Route Route

	app.Post("/users", route.InsertData)
	app.Get("/users", route.GetAllData)
	app.Get("/users/:id_user", route.GetUserByid)

	app.Delete("/users/:id_user", route.Delete)
	app.Put("/users/:id_user", route.Update)

	log.Fatal(app.Listen(":3000"))
}
