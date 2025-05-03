package models

import (
	"time"

	"gorm.io/gorm"
)

type Garden struct{
  ID string `gorm:"primaryKey"`
  UserID string `gorm:"not null;index"`
  User Player `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
  Name string
  Width int `gorm:"default:20"`
  Height int `gorm:"default:20"`
  Theme string `gorm:"default:'default'"`
  Season string `gorm:"default:'spring'"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
