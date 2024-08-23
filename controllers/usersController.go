package controllers

import (
	"artpollybackend/initializers"
	"artpollybackend/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var body struct {
	Name     string
	Email    string
	Password string
}

func Signup(c *gin.Context) {
	if c.ShouldBindJSON(&body) != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Failed to read body",
		}).Error("Signup error")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to hash password")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"error": result.Error.Error(),
		}).Error("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User created successfully")
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	if c.Bind(&body) != nil {
		logrus.WithFields(logrus.Fields{
			"error": "Failed to read body",
		}).Error("Login error")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body.",
		})
		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		logrus.WithFields(logrus.Fields{
			"email": body.Email,
		}).Warn("Invalid email or password")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"email": body.Email,
		}).Warn("Invalid email or password")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 48).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Failed to sign token")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", tokenString, 3600*48, "/", "artpollystudio.ru", true, true)
	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User logged in successfully")

	c.JSON(http.StatusOK, gin.H{
		"userId": user.ID,
		"email":  user.Email,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	if user == nil {
		logrus.Warn("User not found during validation")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.(models.User).ID,
	}).Info("User validated successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func Logout(c *gin.Context) {
	user, _ := c.Get("user")

	c.SetCookie("Authorization", "", -1, "/", "artpollystudio.ru", true, true)

	logrus.WithFields(logrus.Fields{
		"user_id": user.(models.User).ID,
	}).Info("User logged out successfully")

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
		"userId":  user.(models.User).ID,
	})
}

func GetInfo(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		logrus.Warn("Failed to retrieve token from cookie")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized: No token provided",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		logrus.Warn("Invalid token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized: Invalid token",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logrus.Warn("Invalid token claims")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized: Invalid token claims",
		})
		return
	}

	userId, ok := claims["sub"].(float64)
	if !ok {
		logrus.Warn("User ID not found in token claims")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized: User ID not found",
		})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, userId)
	if result.Error != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": userId,
			"error":   result.Error.Error(),
		}).Error("Failed to retrieve user")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user",
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id": user.ID,
	}).Info("User info retrieved successfully")

	c.JSON(http.StatusOK, gin.H{
		"Name":  user.Name,
		"Email": user.Email,
	})
}
