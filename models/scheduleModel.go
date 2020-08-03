package models

import (
	u "../utils"
	"github.com/jinzhu/gorm"
)

// ScheduleModel represents a request to run a command.
type ScheduleModel struct {
	ID           int64  `json:"id"`
	Status       string `form:"status" json:"status"`
	Name         string `form:"name" json:"name"`
	Type         string `form:"type" json:"type"`
	StartDate    string `form:"start_date" json:"start_date"`
	EndDate      string `form:"end_date" json:"end_date"`
	Urlpath      string `form:"url_path" json:"url_path"`
	LastExecute  string `form:"last_execute" json:"last_execute"`
	StatusReturn string `form:"status_return" json:"status_return"`
	Day          string `form:"day" json:"day"`
	Hour         string `form:"hour" json:"hour"`
	Minute       string `form:"minute" json:"minute"`
	User         string `form:"user" json:"user"`
	gorm.Model
}

// TableName Set ScheduleModel table name to be `schedules`.
func (ScheduleModel) TableName() string {
	return "schedules"
}

// CreateSchedule represents a request to run a command.
func (schedule *ScheduleModel) CreateSchedule() map[string]interface{} {
	GetDB().Create(schedule)

	resp := u.Message(true, "success")
	resp["data"] = schedule
	return resp
}
