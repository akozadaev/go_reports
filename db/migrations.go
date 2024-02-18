package database

import (
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&Account{}, &Role{}, &AccountsRoles{}, &Document{}, &Report{}, &ReportStatus{})
	if err != nil {
		log.Printf("[ERROR] failed with %+v", "Migration failed")
	}
	db.Create(&Role{Name: "USER"})
	db.Create(&Role{Name: "DOCTOR"})
	db.Create(&ReportStatus{Name: "NEW"})
	db.Create(&ReportStatus{Name: "PROCESSED"})
	db.Create(&ReportStatus{Name: "ERROR"})
	db.Create(&ReportStatus{Name: "FINISH"})
	db.Migrator().CreateConstraint(&Report{}, "ReportStatus")
	db.Migrator().CreateConstraint(&Report{}, "fk_report_statuses")
	/*	db.Migrator().CreateConstraint(&Account{}, "AccountRoles")
		db.Migrator().CreateConstraint(&Role{}, "AccountRoles")
		db.Migrator().CreateConstraint(&Document{}, "Account")*/
}
