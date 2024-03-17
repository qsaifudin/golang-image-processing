package services

import (
	"bytes"
	"io"

	"github.com/disintegration/imaging"
)

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
