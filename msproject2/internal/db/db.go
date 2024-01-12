package db

import (
	"log"

	"gorm.io/gorm"
   "gorm.io/driver/postgres"
)

func DbConn() *gorm.DB{
  db, err := gorm.Open(
    postgres.Open( "host=localhost user=postgres dbname=msproject2 password=partner port=5432 sslmode=disable"), &gorm.Config{},
  )
  if err != nil {
    log.Fatalf("There was error connecting to the database: %v", err)
  }
  return db
}