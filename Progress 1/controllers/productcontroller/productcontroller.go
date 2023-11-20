package productcontroller

import (
	"TODOLIST/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var Homes []models.Home

	models.DB.Find(&Homes)
	c.JSON(http.StatusOK, gin.H{"Homes": Homes})

}

func Show(c *gin.Context) {

	var Home models.Home
	id := c.Param("id")

	if err := models.DB.First(&Home, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Homes": Home})

}

func Create(c *gin.Context) {

	var Home models.Home

	if err := c.ShouldBindJSON(&Home); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&Home)
	c.JSON(http.StatusOK, gin.H{"Homes": Home})
}

func Update(c *gin.Context) {

	var Home models.Home
	id := c.Param("id")

	if err := c.ShouldBindJSON(&Home); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&Home).Where("id =?", id).Updates(&Home).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cant update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data has successfully updated"})

}

func Delete(c *gin.Context) {

	var Home models.Home

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&Home, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cant delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been deleted"})

}
