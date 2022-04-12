package utils

import (
	"bytes"
	"encoding/base64"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	unbased, _ := base64.StdEncoding.DecodeString(base64Image)
	r := bytes.NewReader(unbased)
	im, err := jpeg.Decode(r)
	if err != nil {
		log.Fatal(err)
	}
    f, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
    _ = jpeg.Encode(f, im, &jpeg.Options{Quality: 75})
}