package sponsor

import (
	"encoding/json"

	jose "gopkg.in/square/go-jose.v2"
)

type NPSigner struct {
	signer jose.Signer
}

func NewNPSigner(privateKey string) (*NPSigner, error) {
	s := NPSigner{}
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS512, Key: []byte(privateKey)}, nil)
	if err != nil {
		return nil, err
	}
	s.signer = signer
	return &s, nil
}

type NPToken struct {
	Sub string `json:"sub"`
}

func (s *NPSigner) Sign(sub string) (string, error) {
	t := NPToken{
		Sub: sub,
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
