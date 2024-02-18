package database

import (
	"encoding/json"
	"gorm.io/gorm"
)

type AccountsRoles struct {
	AccountId uint `json:"account_id,omitempty"`
	RoleId    uint `json:"role_id,omitempty"`
}

type Role struct {
	ID     uint   `gorm:"primarykey" json:"id,omitempty"`
	Name   string `json:"name,omitempty" :"name"`
	RoleId []AccountsRoles
}
type Account struct {
	gorm.Model `json:"gorm_._model,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	MiddleName string `json:"middle_name,omitempty"`
	AccountId  []AccountsRoles
}

type Document struct {
	gorm.Model `json:"gorm_._model"`
	Patient    uint            `json:"patient,omitempty"`
	Doctor     uint            `json:"doctor,omitempty"`
	Document   json.RawMessage `json:"document,omitempty"`
}

type ReportStatus struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"status,omitempty"`
}

type Report struct {
	gorm.Model   `json:"gorm_._model"`
	FileName     string `json:"file_name,omitempty"`
	DownloadLink string `json:"download_link,omitempty"`
	Status       uint
}
