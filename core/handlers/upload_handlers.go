package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
)

func SaveImageToLocal(c echo.Context) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No image file provided"})
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	targetDir := "payment_screenshot"
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create directory"})
	}

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	destinationPath := filepath.Join(targetDir, fileName)

	dst, err := os.Create(destinationPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message":   "File uploaded successfully",
		"file_path": destinationPath,
	})
}