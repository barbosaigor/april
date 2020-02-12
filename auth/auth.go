package auth

import (
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)

type Credentials struct {
	// whitelist contain authorized users.
	whitelist map[string]struct{}
}

func New() *Credentials {
	c := Credentials{}
	c.whitelist = make(map[string]struct{}, 1)
	return &c
}

// EncryptUser takes credentials and convert to SHA256 string
func EncryptUser(username, password string) string {
	key := fmt.Sprintf("%v.%v", username, password)
	tmp := sha256.Sum256([]byte(key))
	eKey := tmp[:]
	return hex.EncodeToString(eKey)
}

// Register add user to an authorised user list
func (c *Credentials) Register(username, password string) {
	user := EncryptUser(username, password)
	c.whitelist[user] = struct{}{}
}

// IsAuthorised verify if its a valid credential
func (c *Credentials) isAuthorised(token string) bool {
	_, ok := c.whitelist[token]
	return ok
}
