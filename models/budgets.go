package models

import (
    "gorm.io/gorm"
    "time"
)

type Budget struct {
    gorm.Model
    UserID uint
    Name   string  `gorm:"size:100"`
    Amount float64
}


