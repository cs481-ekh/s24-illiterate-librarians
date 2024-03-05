package survey

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AfterSemesterSurveyTakenHandler(c *gin.Context) {
	userId := c.Param("user_id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("TODO: make func to check if semester survey for user: %s has been taken for the current semester", userId), //should return a boolean if it was taken or not


		//I found this in the gorm documentation on how to access a database with a table of type "Product"
		//This might be helpfull to understand how to query the database using gorm.

		// var product Product
  		// db.First(&product, 1) // find product with integer primary key
  		// db.First(&product, "code = ?", "D42") // find product with code D42

		// // Update - update product's price to 200
		// db.Model(&product).Update("Price", 200)
		// // Update - update multiple fields
		// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
		// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	  
		// // Delete - delete product
		// db.Delete(&product, 1)
	})
}
