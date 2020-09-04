package models

import (
	"time"

	u "../utils"
	"github.com/jinzhu/gorm"
)

// ScheduleModel represents a request to run a command.
type ScheduleModel struct {
	Status        int        `form:"status" json:"status"`
	Name          string     `form:"name" json:"name"`
	Type          int        `form:"type" json:"type"`
	StartDate     time.Time  `form:"start_date" json:"start_date"`
	EndDate       *time.Time `form:"end_date" json:"end_date"`
	URLPATH       string     `form:"url_path" json:"url_path"`
	LastExecute   *time.Time `form:"last_execute" json:"last_execute"`
	StatusReturn  *int       `form:"status_return" json:"status_return"`
	Day           *int       `form:"day" json:"day"`
	Hour          *int       `form:"hour" json:"hour"`
	Minute        *int       `form:"minute" json:"minute"`
	Counterfailed int        `form:"counterfailed" json:"counterfailed"`
	User          *string    `form:"user" json:"user"`
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
