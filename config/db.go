package config

import (
 
    "log"
   
    "time"

    "Farmako/model"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB


func ConnectDB() {

   dsn:="host=localhost port =5432 user=postgres password=root dbname =testdb sslmode=disable"

    cfg := &gorm.Config{
        Logger: logger.Default.LogMode(logger.Silent),
    }
    db, err := gorm.Open(postgres.Open(dsn), cfg)
    if err != nil {
        log.Fatalf("failed to connect to Postgres: %v", err)
    }

  
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("fail to get generic DB: %v", err)
    }
    sqlDB.SetMaxIdleConns(5)
    sqlDB.SetMaxOpenConns(20)
    sqlDB.SetConnMaxLifetime(time.Hour)

    if err := db.AutoMigrate(&model.Coupon{}); err != nil {
        log.Fatalf("AutoMigrate fail: %v", err)
    }

    DB = db
    log.Println(" Database connected and migrated")
}
