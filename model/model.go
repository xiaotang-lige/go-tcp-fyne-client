package model

import "time"

type TChatModel struct {
	ID          int64     `gorm:"primarykey" json:"id" gorm:"column:id"`
	CreatedTime time.Time `json:"createdTime" gorm:"column:createdTime"`
	UpdatedTime time.Time `json:"updateTime" gorm:"column:updatedTime"`
	DeletedTime time.Time `gorm:"index" json:"deleteTime" gorm:"column:deletedTime"`
}
