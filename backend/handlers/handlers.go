package handlers

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"image-processing-service/services"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers routes with the Gin router
func RegisterRoutes(router *gin.Engine) {
    router.POST("/convert/png-to-jpeg", ConvertPNGToJPEG)
    router.POST("/resize", ResizeImage)
    router.POST("/compress", CompressImage)
    router.Static("/temp_img", "./temp_img")
}

// convertImageToJPEG is a generic function to handle image conversion to JPEG
func convertImageToJPEG(c *gin.Context, converter func(reader io.Reader) ([]byte, error)) {
    // Retrieve the uploaded file
    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "No file uploaded"})
        return
    }

    // Open the uploaded file
    reader, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to open file"})
        return
    }
    defer reader.Close()

    // Perform image conversion
    jpegBytes, err := converter(reader)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
        return
    }

    // Create the folder if it doesn't exist
    targetDir := "temp_img"
    if err := os.MkdirAll(targetDir, 0755); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to create directory"})
        return
    }

    // Generate a unique filename for the processed image
    tempFile := filepath.Join(targetDir, "processed-image.jpeg")

    // Write the processed image bytes to the temporary file
    if err := ioutil.WriteFile(tempFile, jpegBytes, 0644); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Failed to write temporary file"})
        return
    }

    // Return the URL of the temporary file
    c.JSON(http.StatusOK, gin.H{"success": true, "url": "http://localhost:5000/temp_img/processed-image.jpeg"})
}

// ConvertPNGToJPEG handles the conversion of PNG images to JPEG
func ConvertPNGToJPEG(c *gin.Context) {
    convertImageToJPEG(c, services.ConvertPNGToJPEG)
}

// ResizeImage handles image resizing
func ResizeImage(c *gin.Context) {
    width, _ := strconv.Atoi(c.PostForm("width"))
    height, _ := strconv.Atoi(c.PostForm("height"))
    convertImageToJPEG(c, func(reader io.Reader) ([]byte, error) {
        return services.ResizeImage(reader, width, height)
    })
}

// CompressImage handles image compression
func CompressImage(c *gin.Context) {
    quality, _ := strconv.Atoi(c.PostForm("quality"))
    convertImageToJPEG(c, func(reader io.Reader) ([]byte, error) {
        return services.CompressImage(reader, quality)
    })
}
