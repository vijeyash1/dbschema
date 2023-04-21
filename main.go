package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	dsn := "core:core@tcp(127.0.0.1:3306)/core?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.AutoMigrate(
		&User{},
		&Organization{},
		&Service{},
		&Role{},
		&Action{},
		&Features{},
		&OrganizationUserRole{},
		&RoleActionService{},
		&OrganizationFeatureService{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
