# NitroPay Sponsor Library for Go

## Description

Creates a signed token, for passing user identity to sponsor client library.

SiteID and UserID are required, other fields are optional.

```golang
import sponsor "github.com/ggsoftwarellc/nitropay-sponsor-go/v3"

s := sponsor.NewSigner(privateKey)
userInfo := sponsor.UserInfo{
    SiteID: siteID,
    UserID: userID,
    Name: name,
    Email: email,
    Avatar: avatar,
}
token, _ := s.Sign(userInfo)
```

You can use `GetUserSubscription(userID)` to look up subscription info for a user.
