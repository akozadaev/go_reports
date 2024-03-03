package report

import (
	database "akozadaev/go_reports/db"
	"akozadaev/go_reports/db/model"
	"context"
	"gorm.io/gorm"
)

type ReportRepository interface {
	//GetDocument(ctx context.Context, team *model.Document) error
	FindDocumentById(ctx context.Context, id string) model.Document
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (a *reportRepository) FindDocumentById(ctx context.Context, id string) model.Document {
	db := database.FromContext(ctx, a.db)
	var result model.Document
	db.First(&result, "id = ?", id)

	return result
}
