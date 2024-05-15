package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/hakushigo/pa_c_obat/model"
	"strconv"
)

func AddProduk(ctx fiber.Ctx) error {

	var data struct {
		KategoriProduk []int        `json:"kategori_produk"`
		DataProduk     model.Produk `json:"data_produk"`
	}

	err := json.Unmarshal(ctx.Body(), &data)

	produk := data.DataProduk

	// array of object model
	var KategoriProduks []model.KategoriProduk

	// fetch them
	for _, v := range data.KategoriProduk {
		var tempKategoriProduks model.KategoriProduk
		search := db.Find(&tempKategoriProduks, v)

		if search.RowsAffected == 0 {
			return ctx.Status(fiber.StatusRequestedRangeNotSatisfiable).JSON(fiber.Map{
				"status":  fiber.StatusRequestedRangeNotSatisfiable,
				"message": "the kategori with id " + strconv.Itoa(v) + " you find can't be found",
			})
		}

		KategoriProduks = append(KategoriProduks, tempKategoriProduks)
	}

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	produk.KategoriProduk = KategoriProduks
	db.Create(&produk) // save data

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
	})
}

func ListProduk(ctx fiber.Ctx) error {
	var daftarObat []model.Produk
	db.Preload("KategoriProduk").Find(&daftarObat)

	return ctx.Status(200).JSON(daftarObat)
}

func GetProduk(ctx fiber.Ctx) error {

	var dataObat model.Produk

	id, _ := strconv.Atoi(ctx.Params("id"))
	result := db.Preload("KategoriProduk").Find(&dataObat, id)
	if result.RowsAffected <= 0 {
		ctx.Status(404).JSON(fiber.Map{
			"status": 404,
		})
	}

	return ctx.Status(200).JSON(dataObat)
}

func UpdateProduk(ctx fiber.Ctx) error {
	var data model.Produk

	err := json.Unmarshal(ctx.Body(), &data)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	status := db.Find(&model.Produk{}, id).Updates(&data)

	if status.Error != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status": 500,
			"error":  status.Error,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":        200,
		"rows_affected": status.RowsAffected,
	})
}

func DeleteProduk(ctx fiber.Ctx) error {
	var obat model.Produk

	// first fetch the id
	id, _ := strconv.Atoi(ctx.Params("id"))

	// fetch the Obat object
	err := db.First(&obat, id).Error
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// delete the many-to-many association with KategoriProduk
	err = db.Model(&obat).Association("KategoriProduk").Clear()
	if err != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	// delete the Obat
	status := db.Delete(&obat)
	if status.Error != nil {
		return ctx.JSON(fiber.Map{
			"status": 500,
			"error":  err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":        200,
		"rows_affected": status.RowsAffected,
	})
}
