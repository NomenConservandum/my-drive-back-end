package db

import "time"

// TEMPORARY BLOCK START
const UsersNum = 10
const FilesNum = 20

var Array [UsersNum]User

var ArrayJWT [UsersNum]Tokens

var ArrayFiles [UsersNum * FilesNum]Metadata

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
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Owner     string    `json:"owner"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"time"`
}
