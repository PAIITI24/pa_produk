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
			AppName:   "Produk_Apotek_APP",
		})

	// kategori obat
	server.Post("/produk/kategori", controller.AddKategori)
	server.Get("/produk/kategori", controller.ListKategori)
	server.Get("/produk/kategori/:id", controller.GetKategori)
	server.Put("/produk/kategori/:id", controller.UpdateKategori)
	server.Delete("/produk/kategori/:id", controller.DeleteKategori)

	// obat
	server.Post("/produk/", controller.AddProduk)
	server.Get("/produk/", controller.ListProduk)
	server.Get("/produk/:id", controller.GetProduk)
	server.Put("/produk/:id", controller.UpdateProduk)
	server.Delete("/produk/:id", controller.DeleteProduk)

	server.Listen(":3002")
}
