package document

import (
	"math/rand"
	"time"

	"github.com/spacebin-org/curiosity/config"
	"github.com/spacebin-org/curiosity/database"
	"github.com/spacebin-org/curiosity/models"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// CreateID generates a random string of length `length` using the unix timestamp
func CreateID(length int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, length)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// GetDocument retrieves a document record from the database via `id`
func GetDocument(id string) (*models.Document, error) {
	document := models.Document{}
	err := database.DBConn.Where("id = ?", id).First(&document)

	return &document, err.Error
}

// NewDocument creates a new document record in the database
func NewDocument(content string, extension string) (string, error) {
	id := CreateID(config.Config.Documents.IDLength)

	doc := models.Document{
		ID:        id,
		Content:   content,
		Extension: extension,
	}

	// Create new record in database
	res := database.DBConn.Create(&doc)

	return doc.ID, res.Error
}
