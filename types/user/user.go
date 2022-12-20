type UserRoot struct {
  Result Result `json:"result"i`
}

type Result struct {
  User UserInfo `json:"user"`
}

type UserInfo struct {
  FID int64 `json:"fid"`
  Username string `json:"username"`
  DisplayName string `json:"displayName"`
  Pfp Pfp `json:"pfp"`
  FollowerCount int64 `json:followerCount"`
  FollowingCount int64 `json:"followingCount"`
  ReferrerUsername *string `"json:"referrerUsername,omitemp"`
  ViewerContent ViewerContext `json:"viewerContext"`
}

type Pfp Struct {
  URL string `json:"url"`
  Verified bool `json:"verified"`
}

type ViewerContext struct {
  Following bool `json:"following"`
  FollowedBy bool `json:"followedBy"`
  CanSendDirectCasts bool `json:"canSendDirectCasts"`
}
