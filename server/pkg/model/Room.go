package model

import (
	"time"

	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID        		uint    	`json:"id" gorm:"primary_key"`
	CreatedAt 		time.Time	`json:"created_at"`
	UpdatedAt 		time.Time	`json:"updated_at"`
	DeletedAt 		*time.Time	`json:"deleted_at"`

	RoomID 			string		`json:"room_id"`
	MaxPlayer 		int			`json:"max_player"`
	Public 			bool		`json:"public"`
	Host 			string		`json:"host"`
	ExpiredTime 	time.Time	`json:"expired_time"`
}