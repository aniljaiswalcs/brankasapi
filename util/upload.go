package util

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strings"
)

// SaveToDisk - save file to locals
func SaveToDisk(path string, file multipart.File, extension string) (string, error) {

	// Create a temporary file within images directory
	tempFile, err := ioutil.TempFile(path, "upload-*"+extension)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer tempFile.Close()
	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	return tempFile.Name(), nil
}

// GetFileMIMEType - get mime type of file
func GetFileMIMEType(extension string) string {
	extension = strings.ToLower(extension)
	switch extension {
	// image
	case ".bmp":
		return "image/bmp"
	case ".cod":
		return "image/cis-cod"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".ief":
		return "image/ief"
	case ".jpe":
		return "image/jpeg"
	case ".jpeg":
		return "image/jpeg"
	case ".jpg":
		return "image/jpeg"
	case ".jfif":
		return "image/pipeg"
	case ".svg":
		return "image/svg+xml"
	case ".tif":
		return "image/tiff"
	case ".tiff":
		return "image/tiff"
	case ".ras":
		return "image/x-cmu-raster"
	case ".cmx":
		return "image/x-cmx"
	case ".ico":
		return "image/x-icon"
	case ".pnm":
		return "image/x-portable-anymap"
	case ".pbm":
		return "image/x-portable-bitmap"
	case ".pgm":
		return "image/x-portable-graymap"
	case ".ppm":
		return "image/x-portable-pixmap"
	case ".rgb":
		return "image/x-rgb"
	case ".xbm":
		return "image/x-xbitmap"
	case ".xpm":
		return "image/x-xpixmap"
	case ".xwd":
		return "image/x-xwindowdump"
	case ".png":
		return "image/png"
	default:
		return ""
	}
}
