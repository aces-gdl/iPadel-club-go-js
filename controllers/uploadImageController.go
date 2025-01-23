package controllers

import (
	"fmt"
	"io/ioutil"
	"ipadel-club/initializers"
	"ipadel-club/models"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

const UPLOAD_PATH = "./images/"

func UploadImage(c *gin.Context) {
	img, err := imageupload.Process(c.Request, "file")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer imagen...",
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

	fileType := filepath.Ext(img.Filename)

	fileNameBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al generar ID de imagen...",
		})
		return
	}

	fileName := fmt.Sprintf("%s-%s-%s%s", id, name, strings.TrimSpace(string(fileNameBytes)), fileType)

	img.Save(UPLOAD_PATH + fileName)

	switch name {
	case "Event":
		var event models.Event
		results := initializers.DB.Debug().First(&event, id)
		if results.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Error al buscar evento..."})
			return
		}
		event.ImageURL = fileName
		initializers.DB.Save(&event)
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
