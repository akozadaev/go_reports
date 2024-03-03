package server

import (
	"akozadaev/go_reports/db/authentication"
	database "akozadaev/go_reports/db/model"
	"akozadaev/go_reports/db/report"
	"akozadaev/go_reports/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	reportDb report.ReportRepository
	authDB   authentication.AuthRepository
}

func NewHandler(reportDb report.ReportRepository, authDB authentication.AuthRepository) *Handler {
	return &Handler{
		reportDb: reportDb,
		authDB:   authDB,
	}
}
func Start(db *gorm.DB) {
	//dbg = db
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	v1 := router.Group("v1")
	v1.Use()
	{
		v1.GET("report", setReportTask)
		v1.POST("generate", generateData)
	}
	router.Run("localhost:8080")
}

func generateData(context *gin.Context) {
	type CntRequestBody struct {
		Cnt int
	}

	var bodyCnt CntRequestBody
	if err := context.BindJSON(&bodyCnt); err != nil {
		context.Error(err)
	}
	result := bodyCnt.Cnt
	//dbg.AutoMigrate()
	context.IndentedJSON(http.StatusOK, result)
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

func RouteV1(h *Handler, router *gin.Engine) {

	v1 := router.Group("v1")
	var authService authentication.Authentication
	authService = authentication.NewBasicAuth(h.authDB)

	v1.Use(middleware.AuthenticationMiddleware(authService))
	v1.Use()
	{
		v1.GET("report", setReportTask)
	}

}
