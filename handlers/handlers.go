package handlers

import (
	"github.com/gin-gonic/gin"
	"iae.com/smartpower/models"
	"strconv"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok, gin!",
	})
}

//模拟数据库中的数据
var users []models.User
func init() {
	users = append(users, models.User{ID: 1, Username: "user1", Email: "mail1@ae.com"})
	users = append(users, models.User{ID: 2, Username: "user2", Email: "mail2@ae.com"})
	users = append(users, models.User{ID: 3, Username: "user3", Email: "mail3@ae.com"})
}


func GetUserByID(c *gin.Context){
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	// 模拟查询
	var user models.User
	for _, u := range users {
		if u.ID == id {
			user = u
			break
		}
	}

	if user.ID != 0 {
		c.JSON(200, user)
	}else{
		c.JSON(404, gin.H{"error": "User not found"})
	}
}


