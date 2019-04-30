package sponsor

import (
	"encoding/json"
	"testing"

	jose "gopkg.in/square/go-jose.v2"
)

func TestSign(t *testing.T) {
	key := "placeholder"
	userID := "tester"

	s, err := NewNPSigner(key)
	if err != nil {
		t.Error("Unable to create signer: ", err)
	}
	token, err := s.Sign(userID)
	if err != nil {
		t.Error("Error creating token: ", err)
	}

	object, err := jose.ParseSigned(token)
	if err != nil {
		t.Error("Error parsing: ", err)
	}

	payload, err := object.Verify([]byte(key))
	if err != nil {
		t.Error("Unable to verify: ", err)
	}

	nt := NPToken{}
	err = json.Unmarshal(payload, &nt)
	if err != nil {
		t.Error("Unmarshal payload: ", err)
	}

	if nt.Sub != userID {
		t.Error("payload didn't match")
	}
}
