package controllers

import (
	helper "go-todolist/helpers"
	"go-todolist/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDepartemen godoc
//
//	@Router			/Tasklist/GetDepartemen [Get]
func (repository *InitRepo) GetDepartemen(c *gin.Context) {
	var HasilJwt []models.JwtFetch
	// if err := c.ShouldBindJSON(&departemen); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	helper.MasterQuery = models.Query_MasterDept
	errs := helper.MasterExec_Get(repository.DbPg, &HasilJwt)
	if errs != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": errs})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"error": false,
		"data":  HasilJwt,
	})

}
