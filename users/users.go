package users

import (
	"fictitious-bank.com/backend/helpers"
	"fictitious-bank.com/backend/interfaces"
)

func Login(username string, pass string) map[string]interface{}{
	db := helpers.ConnectDB()
	user := &interfaces.User{}
	// RecordNotFound() old implementations
	if err :=  db.Where("username = ?", username).First(&user).Error; err != nil {
		return map[string]interface{}{"message": "user not found"}
	}
	// minuta 6.08 https://www.youtube.com/watch?v=7qMW2LbG3ik&list=PLi21Ag9n6jJJ5bq77cLYpCgOaONcQNqm0&index=2&ab_channel=Duomly
}