package sponsor

import (
	"encoding/json"
	"time"

	jose "gopkg.in/square/go-jose.v2"
)

type Signer struct {
	signer jose.Signer
}

func NewSigner(privateKey string) (*Signer, error) {
	s := Signer{}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS512, Key: []byte(privateKey)}, nil)
	if err != nil {
		return nil, err
	}
	s.signer = signer
	return &s, nil
}

type Token struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
}

type UserInfo struct {
	UserID string
}

func (s *Signer) Sign(u UserInfo) (string, error) {
	t := Token{
		Sub: u.UserID,
		Iat: time.Now().Unix(),
	}

	payload, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	object, err := s.signer.Sign(payload)
	if err != nil {
		return "", err
	}

	return object.CompactSerialize()
}
