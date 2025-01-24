package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

const UPLOAD_PATH = "images/"

func UploadImage(c *gin.Context) {

	img, err := imageupload.Process(c.Request, "file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer imagen..." + err.Error(),
		})
		return
	}

	id := c.Request.FormValue("id")
	name := c.Request.FormValue("name")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer ID de imagen...",
		})
		return
	}

	if name == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer nombre de imagen...",
		})
		return
	}

	if _, err := os.Stat(UPLOAD_PATH); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		c.JSON(http.StatusNotFound, gin.H{
			"error": "path '" + UPLOAD_PATH + "' not found...",
		})
		return
	}

	fileType := filepath.Ext(img.Filename)

	fileName := fmt.Sprintf("%s-%s%s", id, name, fileType)

	result1 := img.Save(UPLOAD_PATH + fileName)
	if result1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fallo al guardar imagen..." + result1.Error()})
		return
	}

	switch name {
	case "Event":
		var event models.Event
		results := initializers.DB.Debug().First(&event, id)
		if results.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Error al buscar evento..."})
			return
		}
		event.ImageURL = fileName
		results = initializers.DB.Save(&event)

		if results.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Error al Actualizar evento..."})
			return
		}
	case "Club":
		var club models.Club
		results := initializers.DB.Debug().First(&club, id)
		if results.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Error al buscar club..."})
			return
		}
		club.ImageURL = fileName
		initializers.DB.Save(&club)
	}

	c.JSON(http.StatusOK, gin.H{"ImageName": fileName})
}

func GetImage(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer ID de imagen...",
		})
		return
	}

	c.Request.Header.Set("Content-Type", "image/png")

	c.File(fmt.Sprintf("%s%s", UPLOAD_PATH, id))

}
func GetImageThumb(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer ID de imagen...",
		})
		return
	}

	c.Request.Header.Set("Content-Type", "image/png")

	c.File(fmt.Sprintf("%s%s-thumb.jpeg", UPLOAD_PATH, id))
}

func ListImages(c *gin.Context) {
	files, err := ioutil.ReadDir(UPLOAD_PATH)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Fallo al leer el directorio de imágenes",
		})
		return
	}

	var images []string
	for _, file := range files {
		if !file.IsDir() {
			images = append(images, file.Name())
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

func DeleteImage(c *gin.Context) {
	fileName := c.Param("fileName")

	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nombre de archivo no proporcionado",
		})
		return
	}

	filePath := filepath.Join(UPLOAD_PATH, fileName)

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Archivo no encontrado",
		})
		return
	}

	// Attempt to remove the file
	err := os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al eliminar el archivo: " + err.Error(),
		})
		return
	}

	// Check if there's a thumbnail and remove it if it exists
	thumbPath := filepath.Join(UPLOAD_PATH, fileName[:len(fileName)-len(filepath.Ext(fileName))]+"-thumb.jpeg")
	if _, err := os.Stat(thumbPath); err == nil {
		os.Remove(thumbPath) // Ignore error as thumbnail might not always exist
	}

	// Update the database if necessary
	// This is a placeholder - you'll need to implement the logic to update the database
	// based on your specific requirements
	// updateDatabase(fileName)

	c.JSON(http.StatusOK, gin.H{
		"message": "Imagen eliminada con éxito",
	})
}
