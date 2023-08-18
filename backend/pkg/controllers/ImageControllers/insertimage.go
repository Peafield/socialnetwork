package imagecontrollers

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
)

func InsertImage(db *sql.DB) {

}

func UploadImage(db *sql.DB, imageData ImageData, imageFile io.Reader) error {
	// Generate a unique filename for the image
	filename := fmt.Sprintf("%d-%s.jpg", time.Now().Unix(), uuid.New().String())
	filepath := path.Join("image-folder", filename) // Change "image-folder" to your desired folder

	// Create the image file on disk
	outFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, imageFile)
	if err != nil {
		return err
	}

	// Insert the file path into the database
	query := "INSERT INTO images (name, filepath) VALUES (?, ?)"
	_, err = db.Exec(query, imageData.Name, filepath)
	if err != nil {
		return err
	}

	return nil
}
