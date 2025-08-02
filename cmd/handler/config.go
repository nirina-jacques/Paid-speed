package handler

import (
	"fmt"
	"log"
	"os"
	"payspeed/cmd/model"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServerConfig struct {
	Db     *gorm.DB
	Router *mux.Router
}

type loadModel struct {
	Models []string
}

func (cf *ServerConfig) Config() {
	var db *gorm.DB
	var err error
	dbUser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbpassword, host, "3306", dbName)
	maxAttempts := 10
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
		if err != nil {
			log.Printf("Attempt %d: Cannot connected database: %v\n", attempts, err)
			time.Sleep(2 * time.Second)
			continue
		}
		d, _ := db.DB()
		err = d.Ping()
		if err != nil {
			log.Printf("Attempt %d: Ping of database failed: %v\n", attempts, err)
			time.Sleep(2 * time.Second)
		} else {
			log.Println("Database connexion success!")
			break
		}
	}

	if err != nil {
		log.Fatalf("Impossible connected of database after %d attempts: %v\n", maxAttempts, err)
	}
	// var load loadModel
	// load.GetModels(db
	cf.Db = db
	cf.InitRoutes()
}
func (l *loadModel) GetModels(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		// &model.Currency{},
		// &model.Transaction{},
		// &model.ExchangeRate{},
		// &model.Wallet{},
	)
}
