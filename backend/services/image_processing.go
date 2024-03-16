package services

import (
	"bytes"
	"io"

	"github.com/disintegration/imaging"
)

// ConvertPNGToJPEG converts PNG image to JPEG and returns the JPEG image bytes
func ConvertPNGToJPEG(reader io.Reader) ([]byte, error) {
	// Open the PNG image from the reader
	img, err := imaging.Decode(reader)
	if err != nil {
		return nil, err
	}

	// Convert the image to JPEG format and return the bytes
	var buf bytes.Buffer
	err = imaging.Encode(&buf, img, imaging.JPEG)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ResizeImage resizes the image according to the specified dimensions and returns the resized image bytes
func ResizeImage(reader io.Reader, width int, height int) ([]byte, error) {
	// Open the image from the reader
	img, err := imaging.Decode(reader)
	if err != nil {
		return nil, err
	}

	// Resize the image
	resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)

	// Encode the resized image to bytes and return
	var buf bytes.Buffer
	err = imaging.Encode(&buf, resizedImg, imaging.JPEG)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// CompressImage compresses the image to reduce file size while maintaining quality and returns the compressed image bytes
func CompressImage(reader io.Reader, quality int) ([]byte, error) {
	// Open the image from the reader
	img, err := imaging.Decode(reader)
	if err != nil {
		return nil, err
	}

	// Encode the image to JPEG with specified quality and return the bytes
	var buf bytes.Buffer
	err = imaging.Encode(&buf, img, imaging.JPEG, imaging.JPEGQuality(quality))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
