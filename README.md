# NitroPay Sponsor Library for Go

## Description

Creates a signed token, for passing user identity to sponsor client library.

```golang
import sponsor "github.com/ggsoftwarellc/nitropay-sponsor-go/v3"

s := sponsor.NewSigner(privateKey)
userInfo := sponsor.UserInfo{
    SiteID: siteID,
    UserID: userID,
}
token, _ := s.Sign(userInfo)
```
