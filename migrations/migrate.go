package migrations

import (
	"fmt"
	"log"
	"os"

	"github.com/SachinDuhan/multiplayer/config"
	"github.com/SachinDuhan/multiplayer/models"
)

func InitDB(){
  if os.Getenv("RUN_MIGRATION") == "true"{
    err := config.DB.AutoMigrate(&models.Player{})
    if err != nil {
      log.Fatal("Unable to perform migrations: ", err)
    }
    err1 := config.DB.AutoMigrate(&models.Garden{})
    if err1 != nil {
      log.Fatal("Unable to perform migrations: ", err)
    }
    err2 := config.DB.AutoMigrate(&models.GardenItem{})
    if err2 != nil {
      log.Fatal("Unable to perform migrations: ", err)
    }
    err3 := config.DB.AutoMigrate(&models.GardenAsset{})
    if err3 != nil {
      log.Fatal("Unable to perform migrations: ", err)
    }
    fmt.Println("migrations done")
  }
}
