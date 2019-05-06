package sponsor

import (
	"fmt"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestSign(t *testing.T) {
	key := "placeholder"
	userID := "tester"
	siteID := "2"

	s := NewSigner(key)

	userInfo := UserInfo{
		SiteID: siteID,
		UserID: userID,
	}

	token, err := s.Sign(userInfo)
	if err != nil {
		t.Error("Error creating token: ", err)
	}

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		t.Error("Error parsing: ", err)
	}
	
	if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
		if claims["sub"] != userID || claims["iss"] != siteID {
			t.Error("payload didn't match")
		}
	} else {
		t.Error("mismatch claims")
	}
}
