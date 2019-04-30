# NitroPay Sponsor Library for Go

## Description

Currently creates a signed token, for passing user identity to sponsor client library.

```golang
	s, _ := NewNPSigner(privateKey)
	token, _ := s.Sign(userID)
```
