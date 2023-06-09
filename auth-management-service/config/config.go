package config

import (
	"asidikfauzi/reservation-of-sport-fields-golang/lib/utils"
	"asidikfauzi/reservation-of-sport-fields-golang/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()
	if err != nil {
		fmt.Println("Error reading .env file")
	}

	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)

	DB, err = gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//InitMigrate()
	//InitSeed()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Roles{})
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Pesans{})
	DB.AutoMigrate(&models.Persons{})
	DB.AutoMigrate(&models.TipeLapangans{})
	DB.AutoMigrate(&models.Lapangans{})
	DB.AutoMigrate(&models.Bookings{})
	DB.AutoMigrate(&models.Ratings{})
}

func InitSeed() error {

	roles := []models.Roles{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "owner"},
		{ID: 3, Name: "user"},
	}

	hashPassword, _ := utils.HashPassword("Password123")

	users := models.Users{
		ID:       utils.Uuid(),
		Username: "admin",
		Email:    "admin@gmail.com",
		Password: hashPassword,
		Role_id:  1,
		IsActive: true,
	}

	DB.Create(roles)
	DB.Create(users)

	return nil
}
