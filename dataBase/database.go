package database

import (
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"FINALTASK-BTPN-SYARIAH/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "ahmajist:ahmaji95tkr(localhost:3306)/MYSQL Model*"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.User{})
}

func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := DB.First(&user, id).Error
	return &user, err
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func UpdateUser(user *models.User) error {
	return DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return DB.Delete(&models.User{}, id).Error
}
