package contacts

import (
	"github.com/jithinjk/plivoapp/common"
)

// Contact contact struct
type Contact struct {
	ID      int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Email   string `gorm:"type:varchar(100);primary_key;unique_index" json:"email"`
	Name    string `gorm:"size:255" json:"name"`
	Phone   string `gorm:"size:15" json:"phone"`
	Address string `gorm:"type:varchar(100)" json:"address"`
}

// AutoMigrate Migrate the schema of database if needed
func AutoMigrate() {
	db := common.GetDB()

	db.AutoMigrate(&Contact{})
}

// FindContact get person details
func FindContact(condition interface{}) (Contact, error) {
	db := common.GetDB()
	var contactModel Contact
	err := db.Limit(10).First(&contactModel, condition).Error
	return contactModel, err
}

// FindAllContacts get person details
func FindAllContacts(offset, limit int) ([]Contact, error) {
	db := common.GetDB()
	var cts []Contact
	err := db.Offset(offset).Limit(limit).Find(&cts).Error
	return cts, err
}

// GetCount get count of contacts
func GetCount() (int, error) {
	var count int
	db := common.GetDB()
	var cts []Contact
	err := db.Find(&cts).Count(&count).Error
	return count, err
}

// GetTotalPageLimitOffset return limit and offset
func GetTotalPageLimitOffset(page, totalCount int) (int, int, int) {
	limit := 2
	offset := (page - 1) * limit
	totalPages := (totalCount / limit) + 1
	return limit, offset, totalPages
}
