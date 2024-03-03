package database

import (
	"akozadaev/go_reports/db/model"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&model.Account{}, &model.Role{}, &model.AccountsRoles{}, &model.Document{}, &model.Report{}, &model.ReportStatus{})
	if err != nil {
		log.Printf("[ERROR] failed with %+v", "Migration failed")
	}
	db.Create(&model.Role{Name: "USER"})
	db.Create(&model.Role{Name: "DOCTOR"})
	db.Create(&model.ReportStatus{Name: "NEW"})
	db.Create(&model.ReportStatus{Name: "PROCESSED"})
	db.Create(&model.ReportStatus{Name: "ERROR"})
	db.Create(&model.ReportStatus{Name: "FINISH"})
	db.Migrator().CreateConstraint(&model.Report{}, "ReportStatus")
	db.Migrator().CreateConstraint(&model.Report{}, "fk_report_statuses")
	/*	db.Migrator().CreateConstraint(&Account{}, "AccountRoles")
		db.Migrator().CreateConstraint(&Role{}, "AccountRoles")
		db.Migrator().CreateConstraint(&Document{}, "Account")*/
}
