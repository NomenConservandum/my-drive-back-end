package utils

import (
	"myDrive/db"
	"net/http"
	"strings"
)

func GetUserFromRequest(r *http.Request) db.User { // a dummy function
	var res db.User
	var JWT = strings.Split(r.Header.Get("Authorization"), " ")[1]

	for i := 0; i < db.UsersNum; i++ {
		if db.Array[i].Username == JWT {
			return db.Array[i]
		}
	}
	return res
}
