package controller

import (
	"fmt"
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {

	src := c.Query("source")

	fmt.Println(src)

	res := []models.SearchRes{}

	err := database.DB.Model(&models.User{}).Preload("Hobbys").Where("name LIKE ?", "%"+src+"%").Find(&res).Error
	if err != nil {
		response.ResponseStatusAccept(c, "search first find error")
		return
	}

	response.ResponseList(c, res)

}
