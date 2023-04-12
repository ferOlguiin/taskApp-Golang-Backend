package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	cookieone, err := c.Request.Cookie("Auth")
	if err != nil {
		panic(err.Error())
	}
	c.SetCookie(cookieone.Name, "", -1, "", "", false, false)

	cookietwo, fail := c.Request.Cookie("CheckAuth")
	if fail != nil {
		panic(err.Error())
	}
	c.SetCookie(cookietwo.Name, "", -1, "", "", false, false)

	c.JSON(http.StatusAccepted, gin.H{"message": "Cookies eliminadas correctamente"})
}
