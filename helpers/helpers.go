package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func HandleErr(err error){
	if err != nil{
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
hased, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
HandleErr(err)
return string(hased)
}



func ConnectDB() *gorm.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/fictitious-bank?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	HandleErr(err)
	return db
}