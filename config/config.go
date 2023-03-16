package config

import (
	"asidikfauzi/reservation-of-sport-fields-golang/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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

	InitMigrate()
	InitSeed()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Pesans{})
	DB.AutoMigrate(&models.Persons{})
	DB.AutoMigrate(&models.TipeLapangans{})
	DB.AutoMigrate(&models.Lapangans{})
	DB.AutoMigrate(&models.Bookings{})
	DB.AutoMigrate(&models.Ratings{})
}

func InitSeed() error {

	uuid := uuid.New().String()
	hashPassword, _ := HashPassword("12345678")

	user := models.Users{
		ID:       uuid,
		Username: "admin",
		Email:    "admin@gmail.com",
		Password: hashPassword,
		Role:     "admin",
		IsActive: true,
	}

	DB.Create(user)

	return nil
}
