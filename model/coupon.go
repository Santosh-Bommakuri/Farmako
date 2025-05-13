package model

import (
    "time"

    "gorm.io/gorm"
)

type Coupon struct {
    ID            uint           `gorm:"primaryKey"`
    Code          string         `gorm:"unique;not null"`
    DiscountType  string
    DiscountValue float64
    LimitPerUser  int
    UsageMode     string
    MinOrderValue float64
    ExpiresAt     time.Time
    CreatedAt     time.Time
    UpdatedAt     time.Time
    DeletedAt     gorm.DeletedAt `gorm:"index"`
}
