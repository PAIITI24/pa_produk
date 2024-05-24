package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hakushigo/pa_c_obat/controller"
	"github.com/hakushigo/pa_c_obat/helper"
)

func main() {

	// Migrate table
	helper.Migrator()

	// declare server
	server := fiber.New(
		fiber.Config{
			Immutable: false,
			AppName:   "barang_apotek_APP",
		})

	// kategori barang
	server.Post("/barang/kategori", controller.AddKategori)
	server.Get("/barang/kategori", controller.ListKategori)
	server.Get("/barang/kategori/:id", controller.GetKategori)
	server.Put("/barang/kategori/:id", controller.UpdateKategori)
	server.Delete("/barang/kategori/:id", controller.DeleteKategori)

	// barang
	server.Post("/barang/", controller.AddBarang)
	server.Get("/barang/", controller.ListBarang)
	server.Get("/barang/:id", controller.GetBarang)
	server.Put("/barang/:id", controller.UpdateBarang)
	server.Delete("/barang/:id", controller.DeleteBarang)

	server.Listen(":3002")
}
