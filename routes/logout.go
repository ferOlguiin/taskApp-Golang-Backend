package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	cookieone, err := c.Request.Cookie("Auth")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("esta es la cookie de auth:", cookieone)
	//cookieone.Name = "Deleted"
	//cookieone.Value = "Unuse"
	//cookieone.Expires = time.Unix(1414414788, 1414414788000)

	cookietwo, fail := c.Request.Cookie("CheckAuth")
	if fail != nil {
		panic(err.Error())
	}
	fmt.Println("esta es la cookie de checkauth", cookietwo)
	//cookietwo.Name = "Borrada"
	//cookietwo.Value = "Sin uso"
	//cookietwo.Expires = time.Unix(1414414788, 1414414788000)

	//c.JSON(http.StatusAccepted, gin.H{"message": "Cookies eliminadas correctamente"})
}
