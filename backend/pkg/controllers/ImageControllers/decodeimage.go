package imagecontrollers

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

func InsertImage(db *sql.DB) {

}

func DecodeImage(imageData string) (string, error) {
	// Extract the Base64 encoded data - removing the data:image/png;base64, part
	base64Data := imageData[strings.Index(imageData, ",")+1:]

	// Decode the Base64 data
	fileData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("invalid image")
	}

	// Generate a unique file name and save the file
	filePath := "pkg/db/images/" + uuid.New().String()
	err = os.WriteFile(filePath, fileData, 0666)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("error saving the file")
	}

	return filePath, nil
}
