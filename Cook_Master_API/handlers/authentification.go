package handlers

import (
	"fmt"
	"gastroguru/database"
	"gastroguru/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	u, err := user.FindUserByEmail(email)
	if err != nil {
		return echo.NewHTTPError(500, "Internal Server Error")
	}
	var last_name, first_name, picture, role, address, phone string
	last_name = u.Last_Name
	first_name = u.First_Name
	picture = u.Picture
	role = u.Role
	address = u.Address
	phone = u.Phone
	token, err := generateAccessToken(last_name, first_name, email, picture, role, address, phone)
	if err != nil {
		fmt.Println("Token generation error:", err)
		return echo.NewHTTPError(500, "Failed to generate token")
	}

	return ctx.JSON(200, map[string]string{
		"token": token,
	})
}

// Fonction pour générer un token d'accès JWT
func generateAccessToken(last_name, first_name, email, picture, role, address, phone string) (string, error) {
	// Création des revendications du token
	claims := jwt.MapClaims{
		"last_name":  last_name,
		"first_name": first_name,
		"email":      email,
		"picture":    picture,
		"role":       role,
		"address":    address,
		"phone":      phone,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Expiration du token après 24 heures
	}

	// Création du token JWT avec les revendications et la signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
