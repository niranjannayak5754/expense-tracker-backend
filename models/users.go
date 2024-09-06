package models

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    gorm.Model
    Name     string `gorm:"size:100"`
    Email    string `gorm:"size:100;unique"`
    Password string `gorm:"size:255"`
}