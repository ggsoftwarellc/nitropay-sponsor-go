# NitroPay Sponsor Library for Go

## Description

Currently creates a signed token, for passing user identity to sponsor client library.

```golang
import sponsor "github.com/ggsoftware/nitropay-sponsor-go/v2"

s := sponsor.NewSigner(privateKey)
userInfo := UserInfo{
    UserID: userID,
}
token, _ := s.Sign(userInfo)
```
