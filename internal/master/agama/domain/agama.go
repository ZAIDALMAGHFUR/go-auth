package domain

import "time"

type Agama struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:100;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (Agama) TableName() string {
    return "agama"
}
