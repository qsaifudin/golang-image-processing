package main_test

import (
	"bytes"
	"image-processing-service/handlers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// performRequest performs an HTTP request to the provided router with the given method, path, and request body
func performRequest(router *gin.Engine, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "multipart/form-data")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func TestConvertPNGToJPEG(t *testing.T) {
    // Initialize Gin router
    gin.SetMode(gin.TestMode)
    router := gin.Default()
    router.POST("/convert/png-to-jpeg", handlers.ConvertPNGToJPEG)

    // Open a test PNG file
    file, err := os.Open("testdata/test.png")
    if err != nil {
        t.Fatal(err)
    }
    defer file.Close()

    // Read the file content to a byte slice
    fileContent, err := ioutil.ReadAll(file)
    if err != nil {
        t.Fatal(err)
    }

    // Create a new reader from the file content
    fileReader := bytes.NewReader(fileContent)

    // Create a new multipart writer
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)

    // Add the form field for image
    part, err := writer.CreateFormFile("image", "test.png")
    if err != nil {
        t.Fatal(err)
    }
    _, err = io.Copy(part, fileReader)
    if err != nil {
        t.Fatal(err)
    }

    // Close the multipart writer to add the boundary parameter
    writer.Close()

    // Perform the HTTP request
    req, err := http.NewRequest("POST", "/convert/png-to-jpeg", body)
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Check the response status code
    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
    }
}


func TestResizeImage(t *testing.T) {
	// Initialize Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/resize", handlers.ResizeImage)

	// Open a test PNG file
	file, err := os.Open("testdata/test.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the form fields
	part, err := writer.CreateFormFile("image", "test.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}

	// Write the additional form fields
	err = writer.WriteField("width", "100")
	if err != nil {
		t.Fatal(err)
	}
	err = writer.WriteField("height", "100")
	if err != nil {
		t.Fatal(err)
	}

	// Close the multipart writer to add the boundary parameter
	writer.Close()

	// Perform the HTTP request
	req, err := http.NewRequest("POST", "/resize", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}

func TestCompressImage(t *testing.T) {
	// Initialize Gin router
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/compress", handlers.CompressImage)

	// Open a test PNG file
	file, err := os.Open("testdata/test.png")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the form field for image
	part, err := writer.CreateFormFile("image", "test.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}

	// Add the form field for quality
	err = writer.WriteField("quality", "80")
	if err != nil {
		t.Fatal(err)
	}

	// Close the multipart writer to add the boundary parameter
	writer.Close()

	// Perform the HTTP request
	req, err := http.NewRequest("POST", "/compress", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}
}
