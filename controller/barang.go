package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/hakushigo/pa_c_obat/model"
	"strconv"
)

func AddBarang(ctx fiber.Ctx) error {

	var data struct {
		KategoriBarang []int        `json:"kategori_barang"`
		DataBarang     model.Barang `json:"data_barang"`
	}

	err := json.Unmarshal(ctx.Body(), &data)

	barang := data.DataBarang

	// array of object model
	var KategoriBarangs []model.KategoriBarang

	// fetch them
	for _, v := range data.KategoriBarang {
		var tempKategoriBarangs model.KategoriBarang
		search := db.Find(&tempKategoriBarangs, v)

		if search.RowsAffected == 0 {
			return ctx.Status(fiber.StatusRequestedRangeNotSatisfiable).JSON(fiber.Map{
				"status":  fiber.StatusRequestedRangeNotSatisfiable,
				"message": "the kategori with id " + strconv.Itoa(v) + " you find can't be found",
			})
		}

		KategoriBarangs = append(KategoriBarangs, tempKategoriBarangs)
	}

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	barang.KategoriBarang = KategoriBarangs
	db.Create(&barang) // save data

	return ctx.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   data,
	})
}

func ListBarang(ctx fiber.Ctx) error {
	var daftarObat []model.Barang
	db.Preload("KategoriBarang").Find(&daftarObat)

	return ctx.Status(200).JSON(daftarObat)
}

func GetBarang(ctx fiber.Ctx) error {

	var dataObat model.Barang

	id, _ := strconv.Atoi(ctx.Params("id"))
	result := db.Preload("KategoriBarang").Find(&dataObat, id)
	if result.RowsAffected <= 0 {
		ctx.Status(404).JSON(fiber.Map{
			"status": 404,
		})
	}

	return ctx.Status(200).JSON(dataObat)
}

func UpdateBarang(ctx fiber.Ctx) error {
	var data model.Barang

	err := json.Unmarshal(ctx.Body(), &data)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
		})
	}

	id, _ := strconv.Atoi(ctx.Params("id"))

	status := db.Find(&model.Barang{}, id).Updates(&data)

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

func DeleteBarang(ctx fiber.Ctx) error {
	var obat model.Barang

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

	// delete the many-to-many association with KategoriBarang
	err = db.Model(&obat).Association("KategoriBarang").Clear()
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
