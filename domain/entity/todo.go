package entity

import "time"

type Todo struct {
	ID 					int
	Title 			string
	Description string
	CreatedAt 	time.Time
	UpdateAt 		time.Time
}

type Todos []*Todo