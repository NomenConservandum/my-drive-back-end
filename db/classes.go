package db

import "time"

// TEMPORARY BLOCK START
const UsersNum = 10

var Array [UsersNum]User

var ArrayJWT [UsersNum]Tokens

// TEMPORARY BLOCK END

type Message struct {
	Message string `json:"message"`
}

type Tokens struct {
	Refresh string `json:"refresh"`
	Access  string `json:"access"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Metadata struct {
	ID        string
	Name      string
	Owner     string
	Size      int64
	CreatedAt time.Time
}
