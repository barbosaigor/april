package auth

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New("my secret")
	if c == nil {
		t.Error("New have not created object")
	}
	if c.key == nil {
		t.Error("New have not created an key")
	}
}

func TestRegister(t *testing.T) {
	c := New("my secret")
	c.Register("bob", "password123")
	c.Register("garry", "password321")
	c.Register("alice", "password213")
	if len(c.login) != 3 {
		t.Errorf("Register expected %v but got %v users", 3, len(c.login))
	}
}

func TestGetToken(t *testing.T) {
	c := New("my secret")
	c.Register("bob", "password123")
	c.Register("garry", "password321")
	c.Register("alice", "password213")

	token, err := c.getToken("garry", "password321")
	if err != nil {
		t.Errorf("getToken got an error: %v", err.Error())
	}
	if token == "" {
		t.Errorf("getToken got an empty token")
	} else {
		t.Logf("getToken generated token: %v", token)
	}

	_, err = c.getToken("garry", "invalid")
	if err != errUnauthorized {
		t.Errorf("getToken fail to authenticate user")
	}

	_, err = c.getToken("invalidUser", "password321")
	if err != errUnauthorized {
		t.Errorf("getToken fail to authenticate user")
	}
}