package models

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Player struct {
  ID          string            `gorm:"primaryKey"`
  Name        string          `gorm:"size:255;not null"`
  Email       string          `gorm:"unique;not null"`
  Password    string          `gorm:"not null"`
  Gardens []Garden `gorm:"foreignKey:UserID"`
  RefreshToken string
  CreatedAt   time.Time
  UpdatedAt   time.Time
  DeletedAt   gorm.DeletedAt  `gorm:"index"` // Soft delete
}

func (u *Player) BeforeCreate(tx *gorm.DB) (err error) {
  u.ID = uuid.NewString()

  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  u.Password = string(hashedPassword)

  return nil
}

func (u *Player) GenerateRefreshToken() (string, error) {
    expiryDays, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY"))
    if err != nil {
        return "", errors.New("Invalid REFRESH_TOKEN_EXPIRY value")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": u.ID,
        "exp":     time.Now().Add(time.Duration(expiryDays) * 24 * time.Hour).Unix(),
    })

    signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func (u *Player) GenerateAccessToken() (string, error) {
    expiryMinutes, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY"))
    if err != nil {
        return "", errors.New("Invalid ACCESS_TOKEN_EXPIRY value")
    }

    fmt.Println("JWT Secret in Go:", os.Getenv("JWT_SECRET"))

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  u.ID,
        "username": u.Name,
        "exp":      time.Now().Add(time.Duration(expiryMinutes) * time.Minute).Unix(),
    })

    signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func (u *Player) CheckPassword(password string) error {
  correct := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
  return correct
}
