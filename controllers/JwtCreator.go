package controllers

import (
	"fmt"
	helper "golang-master-jwt/helpers"
	"golang-master-jwt/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// TaskListJwt godoc
//
//	@Router			/JwtCreator/TaskListJwt [Post]
func (repository *InitRepo) TaskListJwt(c *gin.Context) {
	var HasilJwt models.JwtFetch

	if err := c.ShouldBindJSON(&HasilJwt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//helper.MasterQuery = models.Query_MasterDept
	//errs := helper.MasterExec_Get(repository.DbPg, &HasilJwt)
	//if errs != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs})
	//	return
	//}
	JWT, err := MasterGenerateJWT(HasilJwt.Userstampt)
	if err != nil {
		log.Printf("Failed Creating JWT")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"error": false,
		"data":  JWT,
	})
}

func MasterGenerateJWT(valueofcalling string) (string, error) {
	Jwt_Secret := helper.GodotEnv("JWT_SECRET")
	ISS_KEY := helper.GodotEnv("JWT_ISS")

	if Jwt_Secret == "" || ISS_KEY == "" {
		return "", fmt.Errorf("JWT Secret or Issuer is missing")
	}
	claims := Claims{
		UserID: valueofcalling,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    ISS_KEY,                                             // Issuer
			Subject:   "Authentication",                                    // Subject
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Minute)), // Expiry (5 minutes from now)
			IssuedAt:  jwt.NewNumericDate(time.Now()),                      // Issued time
		},
	}
	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret key
	signedToken, err := token.SignedString([]byte(Jwt_Secret))
	if err != nil {
		return "", fmt.Errorf("error signing the token: %v", err)
	}
	// Output the signed JWT token
	return signedToken, nil
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
