package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetB64Image(image string) string {
	bytes, err := ioutil.ReadFile(image)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += toBase64(bytes)
	return base64Encoding
}

func ToJPG(base64Image string, filePath string) {
	coI := strings.Index(string(base64Image), ",")
	mimeType := strings.TrimSuffix(base64Image[5:coI], ";base64")
	rawImage := string(base64Image)[coI+1:]
	
	unbased, _ := base64.StdEncoding.DecodeString(rawImage)
	r := bytes.NewReader(unbased)

	var im image.Image
	
	switch mimeType {
	case "image/jpeg":
		im, _ = jpeg.Decode(r)
		f, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
    	_ = jpeg.Encode(f, im, &jpeg.Options{Quality: 75})
	case "image/png":
		im, _ = png.Decode(r)
		f, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
    	_ = png.Encode(f, im)
	}
}
