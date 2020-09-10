package models

import (
	"time"

	u "../utils"
	"github.com/jinzhu/gorm"
)

// MessagingModel .
type MessagingModel struct {
	DestinationID string    `form:"destination_id" json:"destination_id"`
	Type          string    `form:"type" json:"type"`
	MessageName   string    `form:"message_name" json:"message_name"`
	MessageTitle  string    `form:"message_title" json:"message_title"`
	MessageBody   string    `form:"message_body" json:"message_body"`
	MessageDesc   *string   `form:"message_desc" json:"message_desc"`
	MessageType   string    `form:"message_type" json:"message_type"`
	ScheduleDate  time.Time `form:"schedule_date" json:"schedule_date"`
	Status        int       `form:"status" json:"status"`
	User          string    `form:"user" json:"user"`
	gorm.Model
}

func (MessagingModel) TableName() string {
	return "messagings"
}

func (messaging *MessagingModel) Validate() (map[string]interface{}, bool) {

	if messaging.DestinationID == "" {
		return u.Message(false, "destination_id required"), false
	}

	if messaging.Type == "" {
		return u.Message(false, "type required"), false
	}

	if messaging.MessageTitle == "" {
		return u.Message(false, "message_title required"), false
	}

	if messaging.MessageBody == "" {
		return u.Message(false, "message_body required"), false
	}

	if messaging.MessageType == "" {
		return u.Message(false, "message_title required"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

// CreateMessaging .
func (messaging *MessagingModel) CreateMessaging() map[string]interface{} {

	if resp, ok := messaging.Validate(); !ok {
		return resp
	}

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
