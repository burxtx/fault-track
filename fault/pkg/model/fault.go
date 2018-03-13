package model

import "time"

type Fault struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	UserName    string    `json:"user_name"`
	State       State     `json:"state"`
	Description string    `json:"description"`
	Hostname    string    `json:"hostname"`
	CreateTime  time.Time `json:"create_time"`
	Category    string    `json:"category"`
	Detail      string    `json:"detail"`
	Comment     string    `json:"comment"`
	UpdateTime  time.Time `xorm:"updated"`
	SaId        int64     `json:"sa_id"`
	SaName      string    `json:"sa_name"`
	Sn          string    `json:"sn"`
	Version     int       `xorm:"version"`
	ParentId    int       `json:"parent_id"`
}

type State int

const (
	Opened State = iota
	Diagnosed
	Informed
	Confirmed
	Closed
	Cancelled
)
