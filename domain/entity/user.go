package entity

import "time"

type User struct {
	ID 				int
	Name 			string
	Email 		string
	PassWord 	string
	CreatedAt time.Time
	TodoIds 	[]int
}