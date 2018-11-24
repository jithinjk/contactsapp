package contacts

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jithinjk/plivoapp/common"
)

// GetContact get contact
func GetContact(c *gin.Context, contactID string) {
	id, err := strconv.Atoi(contactID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No contact found!"})
		return
	}
	contactModel, err := FindContact(&Contact{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No contact found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": contactModel})
}

// GetAllContacts get people details
func GetAllContacts(c *gin.Context) {
	reqPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Invalid Page!"})
		c.AbortWithStatus(404)
		return
	}

	totalCount, cerr := GetCount()
	if cerr != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Count Read Error!"})
		c.AbortWithStatus(404)
		return
	}
	limit, offset, totalPages := GetTotalPageLimitOffset(reqPage, totalCount)
	if reqPage > totalPages {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Page does not exist!"})
		c.AbortWithStatus(404)
		return
	}

	var cts []Contact
	cts, ferr := FindAllContacts(offset, limit)
	if ferr != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No accounts found!"})
		c.AbortWithStatus(404)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": cts, "requestedPage": reqPage, "totalCount": totalCount, "pages": totalPages, "limit": limit, "offset": offset})
}

// CreateContact creates contact
func CreateContact(c *gin.Context) {
	db := common.GetDB()

	var contact Contact

	err := c.BindJSON(&contact)
	if err != nil {
		log.Fatal(err)
		c.Abort()
		return
	}

	if err := db.Where("email = ?", contact.Email).First(&contact).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Conflict. Item Exists"})
		c.Abort()
		return
	}

	if err := db.Create(&contact).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "Conflict:" + err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Contact added successfully"})
}

// UpdateContact updates the contact given by ID
func UpdateContact(c *gin.Context) {
	db := common.GetDB()
	var contact Contact
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&contact).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No account found!"})
		c.AbortWithStatus(404)
		return
	}
	err := c.BindJSON(&contact)
	if err != nil {
		log.Fatal(err)
		c.Abort()
	}

	db.Save(&contact)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Contact updated successfully!"})
}

// DeleteContact deletes the contact corresponding to the given ID
func DeleteContact(c *gin.Context) {
	db := common.GetDB()
	var contact Contact
	contactID := c.Params.ByName("id")

	if err := db.First(&contact, contactID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No account found!"})
		c.AbortWithStatus(404)
		return
	}

	db.Delete(&contact)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Contact" + contactID + " deleted successfully!"})
}
