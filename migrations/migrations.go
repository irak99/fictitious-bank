package migrations

import (
	"fictitious-bank.com/backend/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct{
gorm.Model
Username string
Email string
Password string
}

type Account struct{
	gorm.Model
	Type string
	Name string
	Balance uint
	UserID uint
}

func connectDB() *gorm.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/fictitious-bank?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helpers.HandleErr(err)
	return db
}

func createAccount(){
	db := connectDB()
	users := [2]User{
		{Username: "Nikola", Email: "nikola@nikola.com"},
		{Username: "Kiko", Email: "kiko@kiko.com"},
	}
	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)
		account := Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(1000*int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	db:= connectDB()
	db.AutoMigrate(&User{}, &Account{})
	createAccount()
}