package sponsor

import (
	"time" 

	jwt "github.com/dgrijalva/jwt-go"
)

type Signer struct {
	privateKey []byte
	token *jwt.Token
}

func NewSigner(privateKey string) *Signer {
	s := Signer{}
	s.privateKey = []byte(privateKey)
	s.token = jwt.New(jwt.SigningMethodHS512)

	return &s
}

type UserInfo struct {
	UserID string
}

func (s *Signer) Sign(u UserInfo) (string, error) {
	s.token.Claims = jwt.MapClaims{
		"sub": u.UserID,
		"iat": time.Now().Unix(),
	}
	return s.token.SignedString(s.privateKey)
}
