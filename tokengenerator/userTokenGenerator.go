package tokengenerator

import (
	"golangGinMongo/model"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func GenerateToken(user model.User) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("El archivo .env no se encuentra")
	}
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("No se encontro la llave válida")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Name":     user.Name,
		"Email":    user.Email,
		"Password": user.Password,
	})
	return token.SignedString([]byte(secretKey))

}
