package auth

import (
	"myDrive/db"
	"strings"
)

// NOT REAL JWT YET

// FOR TESTING PURPOSES ONLY

func GenerateJWT(u db.User) db.Tokens {
	var tokens db.Tokens
	tokens.Access = u.Username
	tokens.Refresh = u.Password
	return tokens
}

// requires a string WITH 'Bearer'
func IsValidAccess(jwt string) bool {
	var newstring = strings.Split(jwt, " ")[1]
	for i := 0; i < db.UsersNum; i++ {
		if db.ArrayJWT[i].Access == newstring {
			return true
		}
	}
	return false
}

// requires a string WITHOUT 'Bearer'
func IsValidRefresh(refreshToken string) (bool, db.User) {
	var user db.User
	for i := 0; i < db.UsersNum; i++ {
		if db.ArrayJWT[i].Refresh == refreshToken && refreshToken != "" {
			return true, db.Array[i]
		}
	}
	return false, user
}
