package handlers

import (
	"github.com/ectrc/snow/aid"
	"github.com/gofiber/fiber/v2"
)

func GetCloudStorageFiles(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON([]aid.JSON{})
}

func GetCloudStorageConfig(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(aid.JSON{
		"enumerateFilesPath": "/api/cloudstorage/system",
		"enableMigration": true,
		"enableWrites": true,
		"epicAppName": "Live",
		"isAuthenticated": true,
		"disableV2": true,
		"lastUpdated": "0000-00-00T00:00:00.000Z",
		"transports": []string{},
	})
}

func GetCloudStorageFile(c *fiber.Ctx) error {
	fileName := c.Params("fileName")
	if fileName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(aid.ErrorBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(aid.JSON{})
}

func GetUserStorageFiles(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON([]aid.JSON{})
}

func GetUserStorageFile(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(aid.JSON{})
}

func PutUserStorageFile(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(aid.JSON{})
}