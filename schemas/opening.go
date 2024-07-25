package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
	Role      string     `json:"role"`
	Company   string     `json:"company"`
	Location  string     `json:"location"`
	Remote    bool       `json:"remote"`
	Link      string     `json:"link"`
	Salary    int64      `json:"salary"`
}

func (Opening) TableName() string {
	return "create_opening_requests"
}
