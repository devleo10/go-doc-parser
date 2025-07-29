package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/upload-report", UploadReportHandler)
	r.GET("/dashboard/:userID", DashboardHandler)
	r.GET("/transactions/:userID", TransactionsHandler)
	r.GET("/asset-allocation/:userID", AssetAllocationHandler)
}

func UploadReportHandler(c *gin.Context) {
	// TODO: Implement PDF upload and parsing
	c.JSON(http.StatusOK, gin.H{"message": "Upload endpoint stub"})
}

func DashboardHandler(c *gin.Context) {
	// TODO: Return dashboard data
	c.JSON(http.StatusOK, gin.H{"message": "Dashboard endpoint stub"})
}

func TransactionsHandler(c *gin.Context) {
	// TODO: Return transactions data
	c.JSON(http.StatusOK, gin.H{"message": "Transactions endpoint stub"})
}

func AssetAllocationHandler(c *gin.Context) {
	// TODO: Return asset allocation data
	c.JSON(http.StatusOK, gin.H{"message": "Asset allocation endpoint stub"})
}
