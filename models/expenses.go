package models

import (
    "gorm.io/gorm"
    "time"
)

type Expense struct {
    gorm.Model
    BudgetID    uint
    Description string `gorm:"size:255"`
    Amount      float64
    Date        time.Time
}