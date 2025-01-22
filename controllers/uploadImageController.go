package controllers

import (
	"fmt"
	"net/http"
	"os/exec"
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

	if id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer ID de imagen...",
		})
		return
	}

	fileType := strings.Split(img.ContentType, "/")[1]

	fileNameBytes, err := exec.Command("uuidgen").Output()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al generar ID de imagen...",
		})
		return
	}

	fileName := fmt.Sprintf("%s.%s", strings.TrimSpace(string(fileNameBytes)), fileType)
	thumbName := fmt.Sprintf("%s-thumb.%s", strings.TrimSpace(string(fileNameBytes)), fileType)

	regImage, err := imageupload.ThumbnailJPEG(img, 400, 400, 100)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al crear imagen...",
		})
		return
	}
	regImage.Save(UPLOAD_PATH + fileName)

	thumb, err := imageupload.ThumbnailJPEG(img, 100, 100, 100)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al crear miniatura de imagen...",
		})
		return
	}
	thumb.Save(UPLOAD_PATH + thumbName)

	c.JSON(http.StatusOK, gin.H{"ImageName": fileName, "ThumbName": thumbName})
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
	if c.Request.FormValue("thumb") != "" {
		c.File(fmt.Sprintf("%s%s-thumb.png", UPLOAD_PATH, id))

	} else {
		c.File(fmt.Sprintf("%s%s.jpeg", UPLOAD_PATH, id))
	}
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
