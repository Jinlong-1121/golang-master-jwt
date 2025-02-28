package controllers

import (
	"fmt"
	helper "golang-master-jwt/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// TaskListJwt godoc
//
//	@Router			/JwtCreator/TaskListJwt [Get]
func (repository *InitRepo) TaskListJwt(c *gin.Context) {
	//var HasilJwt []models.JwtFetch

	// if err := c.ShouldBindJSON(&departemen); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	//helper.MasterQuery = models.Query_MasterDept
	//errs := helper.MasterExec_Get(repository.DbPg, &HasilJwt)
	//if errs != nil {
	//	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs})
	//	return
	//}
	JWT, err := MasterGenerateJWT("testing")
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

	Jwt_Secret := []byte(helper.GodotEnv("JWT_SECRET"))
	ISS_KEY := helper.GodotEnv("JWT_ISS")

	if Jwt_Secret == nil && ISS_KEY == "" {
		return "", fmt.Errorf("Jwt Secret is null")
	} else {
		claims := Claims{
			UserID: valueofcalling, // Add your custom claims here
			RegisteredClaims: jwt.RegisteredClaims{
				// Set standard JWT claims
				Issuer:    ISS_KEY,                                             // Issuer
				Subject:   "Authentication",                                    // Subject
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)), // Expiry (5 minutes from now)
				IssuedAt:  jwt.NewNumericDate(time.Now()),                      // Issued time
			},
		}
		// Create a new token with the claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign the token with the secret key
		signedToken, err := token.SignedString(Jwt_Secret)
		if err != nil {
			log.Fatal(err)
		}

		// Output the signed JWT token
		return signedToken, nil
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
