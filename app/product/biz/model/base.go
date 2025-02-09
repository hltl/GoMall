package model

import "time"

type Base struct {
	ID       int `grom:"primary_key"`
	CreateAt time.Time
	UpdateAt time.Time
}
