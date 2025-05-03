package models

import "time"

type GardenAsset struct{
  ID string `gorm:"primaryKey"`
  DisplayName string
  AssetPath string
  CreatedAt time.Time
  UpdatedAt time.Time
}
