package tokengenerator

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func VerifyToken(tokenString string) (jwt.Claims, error) {

	// if err := godotenv.Load(); err != nil {
	// 	log.Println("El archivo .env no se encuentra")
	// }
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("No se encontro la llave válida")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	claims := token.Claims.(jwt.MapClaims)

	return claims, err

}
