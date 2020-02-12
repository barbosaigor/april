package auth

import (
	"encoding/json"
	"net/http"
	"errors"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var errUnauthorized = errors.New("Not authorised user")

// getToken generates a token
func (c *credentials) getToken(username, password string) (string, error) {
	if !c.IsAuthorised(username, password) {
		return "", errUnauthorized
	}
	tk := jwt.New(jwt.SigningMethodHS256)
	token, err := tk.SignedString(c.key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *credentials) MwGenerateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var creds login
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := c.getToken(creds.Username, creds.Password)
		if err == errUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
		})
		next.ServeHTTP(w, r)
	})
}

func (c *credentials) MwAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ck, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := ck.Value
		tkn, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
			return c.key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}