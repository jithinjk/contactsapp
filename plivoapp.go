package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jithinjk/plivoapp/common"
	"github.com/jithinjk/plivoapp/contacts"
)

// GetHandler handler for GET calls
func GetHandler(c *gin.Context) {
	path1 := c.Param("path1")
	path2 := c.Param("path2")

	if path1 == "all" && path2 == "" {
		contacts.GetAllContacts(c)
	} else if path1 != "" && path2 == "details" {
		contactID := path1
		contacts.GetContact(c, contactID)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No contact found. Incorrect Format."})
		c.Abort()
	}
}

// Migrate migrate schema
func Migrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&contacts.Contact{})
}

func main() {
	// open a db connection
	db := common.Init()
	defer db.Close()

	log.Println("Connection Established...")

	db.SingularTable(true)

	//Drops table if already exists
	// db.Debug().DropTableIfExists(&Contact{})

	//Auto create table based on Model
	Migrate(db)

	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("/api/v1/")

	v1.Use(contacts.GetRequestID())

	v1.Use(gin.BasicAuth(gin.Accounts{
		"user1": "hello",
		"user2": "world",
		"user3": "gopher",
	}))
	{
		v1.GET("/contacts/:path1", GetHandler)        //      /v1/contacts/all
		v1.GET("/contacts/:path1/:path2", GetHandler) //      /v1/contacts/<id>/details
		v1.POST("/create", contacts.CreateContact)
		v1.PUT("/update/:id", contacts.UpdateContact)
		v1.DELETE("/delete/:id", contacts.DeleteContact)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": http.StatusNotFound, "message": "Page not found"})
	})

	return router
}

// search by name and email
