package auth

import (
	_ "github.com/dgrijalva/jwt-go"
)

type credentials struct {
	// login contain authorised users.
	// Username as key and password as value
	login map[string]string
	key []byte
}

func New(key string) *credentials {
	c := &credentials{}
	c.login = make(map[string]string, 1)
	c.key = []byte(key)
	return c
}

// Register add user to an authorised user list
func (c *credentials) Register(username, password string) {
	c.login[username] = password
}

// IsAuthorised verify if its a valid credential
func (c *credentials) IsAuthorised(username, password string) bool {
	pwd, ok := c.login[username]
	return ok && password == pwd
}

