package database

import (
	"fmt"
	"log"

	"github.com/xbklyn/getgoal-app/config"
	"github.com/xbklyn/getgoal-app/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	Db *gorm.DB
}

// Migrate implements Database.
func (p *postgresDB) Migrate() {
	migrations.LabelMigrate(p.Db)
}

func (p *postgresDB) GetDb() *gorm.DB {
	return p.Db
}

func NewPostgresDB(cfg *config.Config) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
		cfg.Db.SSLMode,
		cfg.Db.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Default().Fatalf("Fatal error when loading database: %s \n", err)
	}

	return &postgresDB{Db: db}
}
