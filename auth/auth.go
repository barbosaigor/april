package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Credentials can be used for registering and validating users
type Credentials struct {
	// whitelist contain authorized users.
	whitelist map[string]struct{}
}

// New creates a Crendentials structure
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

// Remove deletes an user from authorised user list
func (c *Credentials) Remove(username string) {
	delete(c.whitelist, username)
}

// IsAuthorised verify if its a valid credential
func (c *Credentials) IsAuthorised(token string) bool {
	_, ok := c.whitelist[token]
	return ok
}
