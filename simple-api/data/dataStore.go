package data

import (
	"alvinr.ca/learn-go/api/models"
)

// BooksDB to simulate persistent storage
var BooksDB []models.Book

// PopulateDatastore with books
func PopulateDatastore() {
	BooksDB = append(BooksDB,
		models.Book{ID: "1", ISBN: "9781101965337", Title: "Age of Myth", Author: &models.Author{FirstName: "Michael J.", LastName: "Sullivan"}},
		models.Book{ID: "2", ISBN: "9781101965368", Title: "Age of Swords", Author: &models.Author{FirstName: "Michael J.", LastName: "Sullivan"}},
		models.Book{ID: "3", ISBN: "9781101965399", Title: "Age of War", Author: &models.Author{FirstName: "Michael J.", LastName: "Sullivan"}},
		models.Book{ID: "4", ISBN: "9781944145293", Title: "Age of Legend", Author: &models.Author{FirstName: "Michael J.", LastName: "Sullivan"}},
		models.Book{ID: "5", ISBN: "9781944145392", Title: "Age of Death", Author: &models.Author{FirstName: "Michael J.", LastName: "Sullivan"}},
	)
}
