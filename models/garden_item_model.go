package models

import (
	"time"

	"gorm.io/datatypes"
)

type GardenItem struct {
  ID uint `gorm:"primaryKey;autoIncrement"`
  GardenID string `gorm:"not null"`
  AssetID string `gorm:"not null"`
  X int `gorm:"not null"`
  Y int `gorm:"not null"`
  ZIndex int `gorm:"default:0"`
  Rotation float64 `gorm:"default:0"`
  ScaleX float64 `gorm:"default:1"`
  ScaleY float64 `gorm:"default:1"`
  Properties datatypes.JSON
  CreatedAt time.Time
  UpdatedAt time.Time
}
