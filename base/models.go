package base


import (
  "time"
)


type Model struct {
  ID        uint      `gorm:"primary_key"`
  CreatedAt time.Time `gorm:"type:timestamp with time zone;default:current_timestamp"`
  UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:current_timestamp"`
  DeletedAt time.Time
}
