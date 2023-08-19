package model

import (
	"time"
)

type User struct {
	ID        	uint    	`json:"id" gorm:"primary_key"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`
	DeletedAt 	*time.Time	`json:"deleted_at"`

	Username 	string		`json:"username"`
	Password 	string		`json:"password"`
	Email 		string		`json:"email"`
	Avatar 		string		`json:"avatar"`
	Oauth 		bool		`json:"oauth"`
}