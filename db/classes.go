package db

// TEMPORARY BLOCK START
var Array [10]User

// TEMPORARY BLOCK END

type Error struct {
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
	//
}
