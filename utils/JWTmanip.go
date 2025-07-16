package utils

import (
	"myDrive/db"
	"net/http"
)

func GetUserFromRequest(r *http.Request) db.User { // a dummy function
	var res db.User
	res.Username = "a"
	res.Password = "a"
	return res
}
