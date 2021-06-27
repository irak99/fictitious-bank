package helpers

import (
	"golang.org/x/crypto/bcrypt"
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