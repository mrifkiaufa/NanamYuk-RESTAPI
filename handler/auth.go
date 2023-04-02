package handler

import (
	"nanam-yuk/auth"
	"nanam-yuk/initializers"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Email string
		Password string
		Name string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Failed to read body",
		})
		return
	}

	if (body.Name == "" && body.Email == "" && body.Password == "") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Please fil the data first",
		})
		return
	}

	if (body.Name == "" || body.Name == " ") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Please fil the name first",
		})
		return
	}

	if (body.Email == "" || body.Email == " ") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Please fil the email first",
		})
		return
	}

	if (body.Password == "" || body.Password == " ") {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Please fil the password first",
		})
		return
	}

	if (len(body.Password) < 6) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Password cannot shorter than 6 characters",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Failed to hash password",
		})
		return
	}

	user := auth.Auth{Email: body.Email, Password: string(hash), Name: body.Name}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Failed to create Account",
		})
		return
	}

	userResponse := convertToAuthResponse(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24  * 7).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*7, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user": userResponse,
		"userId": user.ID,
		"token": tokenString,
		"message": "Register successfully",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Failed to read body",
		})
		return
	}

	var user auth.Auth
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Email tidak tersedia",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Password yang anda masukkan salah",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24  * 7).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "Failed to create token",
		})
		return
	}

	userResponse := convertToAuthResponse(user)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*7, "", "", false, true)
	
	c.JSON(http.StatusOK, gin.H{
		"user": userResponse,
		"userId": user.ID,
		"token": tokenString,
		"message": "Login successfully",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "Logout successfully",
	})
}

func convertToAuthResponse(a auth.Auth) auth.AuthResponse {
	return auth.AuthResponse{
		Email: a.Email,
		Password: a.Password,
		Name: a.Name,
	}
}