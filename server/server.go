package server

import (
	database "akozadaev/go_reports/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var dbg *gorm.DB

func Start(db *gorm.DB) {
	//dbg = db
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	v1 := router.Group("v1")
	v1.Use()
	{
		v1.GET("report", setReportTask)
	}
	router.Run("localhost:8080")
}

func setReportTask(context *gin.Context) {
	raw := &database.Report{
		FileName:     "report",
		DownloadLink: "http://report.csv",
		Status:       1,
	}
	// TODO insert
	context.IndentedJSON(http.StatusCreated, raw)
}
