package sponsor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Signer struct {
	privateKey []byte
	token      *jwt.Token
}

func NewSigner(privateKey string) *Signer {
	s := Signer{}
	s.privateKey = []byte(privateKey)
	s.token = jwt.New(jwt.SigningMethodHS512)

	return &s
}

type UserInfo struct {
	SiteID string
	UserID string
	Name   string
	Email  string
	Avatar string
}

func (s *Signer) Sign(u UserInfo) (string, error) {
	s.token.Claims = jwt.MapClaims{
		"iss":    u.SiteID,
		"sub":    u.UserID,
		"iat":    time.Now().Unix(),
		"name":   u.Name,
		"email":  u.Email,
		"avatar": u.Avatar,
	}
	return s.token.SignedString(s.privateKey)
}

type SubscriptionInfo struct {
	Tier struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Order       int    `json:"order"`

		Benefits []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"benefits"`
	} `json:"tier"`
	Status          string `json:"status"`
	SubscribedUntil string `json:"subscribedUntil"`
}

func (s *Signer) GetUserSubscription(userID string) (*SubscriptionInfo, error) {
	transport := &http.Transport{
		TLSHandshakeTimeout:   15 * time.Second,
		ResponseHeaderTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   20 * time.Second,
	}

	url := fmt.Sprintf("https://sponsor-api.nitropay.com/v1/users/%s/subscription", url.PathEscape(userID))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", string(s.privateKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	si := SubscriptionInfo{}

	err = json.Unmarshal(body, &si)
	if err != nil {
		return nil, err
	}

	return &si, err
}
