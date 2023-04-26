package app

import (
	"github.com/gin-gonic/gin"
	"bird_golang_back/internal/db"
	"strings"
	"net/http"
)
func UpdateBirdData (c *gin.Context) {
	var information BirdData
	if err := c.ShouldBind(&information); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error Bad request"})
		return
	}
	err := db.InsertBirdData(information.Idx,information.Name,information.Description)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"Update Bird data success"})
}
func UpdateBirdDes (c *gin.Context) {
	var information BirdInsertDes
	if err := c.ShouldBind(&information); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error Bad request"})
		return
	}

	err := db.InsertBirdDescription(information.Idx,information.Description)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message":"Update Bird description success"})
}

func GetBirdData (c *gin.Context) {
	var searchKey BirdSearchKey
	if err := c.ShouldBind(&searchKey); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Error Bad request"})
		return
	}
	key := *searchKey.Key
	key = strings.ReplaceAll(key,"%20"," ")
	b,err := db.SearchBird(*searchKey.Key)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{ "bird": b})

}
