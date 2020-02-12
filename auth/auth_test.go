package auth

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	if c == nil {
		t.Error("New: have not created object")
	}
	if c.whitelist == nil {
		t.Error("New: have not created whitelist")
	}
}

func TestRegister(t *testing.T) {
	c := New()
	c.Register("bob", "password123")
	c.Register("garry", "password321")
	c.Register("alice", "password213")
	if len(c.whitelist) != 3 {
		t.Errorf("Register: expected %v but got %v users", 3, len(c.whitelist))
	}
}

func TestEncryptUser(t *testing.T) {
	t.Log("EncryptUser: token ", EncryptUser("bob", "password123"))
}
