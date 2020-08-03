package models

import (
	"time"

	u "../utils"
	"github.com/jinzhu/gorm"
)

// MessagingModel .
type MessagingModel struct {
	ID            int64     `json:"id"`
	DestinationID string    `form:"destination_id" json:"destination_id"`
	Type          string    `form:"type" json:"type"`
	MessageName   string    `form:"message_name" json:"message_name"`
	MessageTitle  string    `form:"message_title" json:"message_title"`
	MessageBody   string    `form:"message_body" json:"message_body"`
	MessageDesc   string    `form:"message_desc" json:"message_desc"`
	MessageType   string    `form:"message_type" json:"message_type"`
	ScheduleDate  time.Time `form:"schedule_date" json:"schedule_date"`
	Status        string    `form:"status" json:"status"`
	User          string    `form:"user" json:"user"`
	gorm.Model
}

func (MessagingModel) TableName() string {
	return "messagings"
}

// CreateMessaging .
func (messaging *MessagingModel) CreateMessaging() map[string]interface{} {

	GetDB().Create(messaging)

	resp := u.Message(true, "success")
	resp["data"] = messaging
	return resp
}

// GetMessaging .
func GetMessaging(id int) *MessagingModel {
	messaging := &MessagingModel{}

	err := GetDB().First(&messaging, id).Error
	if err != nil {
		return nil
	}

	return messaging

}
