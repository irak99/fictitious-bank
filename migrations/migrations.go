package migrations

import (
	"fictitious-bank.com/backend/helpers"
	"fictitious-bank.com/backend/interfaces"
)




func createAccount(){
	db := helpers.ConnectDB()
	users := &[2]interfaces.User{
		{Username: "Nikola", Email: "nikola@nikola.com"},
		{Username: "Kiko", Email: "kiko@kiko.com"},
	}
	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)
		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(1000*int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db:= helpers.ConnectDB()
	db.AutoMigrate(User, Account)
	createAccount()
}