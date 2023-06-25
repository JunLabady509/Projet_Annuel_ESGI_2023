package handlers

import (
	"fmt"
	"gastroguru/database"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Clé secrète utilisée pour signer les tokens JWT
var jwtSecret = []byte("monTokenSuperSafe")

func auth(email, password string, ctx echo.Context) error {
	// Vérifie si l'adresse email existe dans la base de données et qu'il correspond au mot de passe.
	exists, err := database.UserExists(email)
	if err != nil {
		fmt.Println("Database error:", err)
		return echo.NewHTTPError(500, "Internal Server Error")
	}

	if !exists {
		fmt.Println("Email doesn't exist")
		return echo.NewHTTPError(401, "Email doesn't exist")
	}

	// Vérifie si le mot de passe correspond à l'adresse email
	match, err := database.PasswordMatchesEmail(password, email)
	if err != nil {
		fmt.Println("Database error:", err)
		return echo.NewHTTPError(500, "Internal Server Error")
	}

	if !match {
		return echo.NewHTTPError(401, "Password doesn't match")
	}

	// Génère un token d'accès
	token, err := generateAccessToken(email)
	if err != nil {
		fmt.Println("Token generation error:", err)
		return echo.NewHTTPError(500, "Failed to generate token")
	}

	return ctx.JSON(200, map[string]string{
		"token": token,
	})
}

// Fonction pour générer un token d'accès JWT
func generateAccessToken(email string) (string, error) {
	// Création des revendications du token
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expiration du token après 24 heures
	}

	// Création du token JWT avec les revendications et la signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
