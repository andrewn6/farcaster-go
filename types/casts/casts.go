package casts

type CastRoot struct {
  Result Result `json:"result"`
  Next Next `json:"next"`
}

type Result struct {
  Casts []Cast `json:"casts"`
}

type Cast struct {
  Hash string `json:"hash"`
  ThreadHash string `json:"threadHash"`
  Author Author `json:"author"`
  Timestamp int64 `json:"timestamp"`
  Replies Replies `json:"replies"`
  Reactions Reactions `json:"reactions"`
  Recasts Recasts `json:"recasts"`
  Watches Watches `json:"watches"`
  Recast bool `json:"recast"`
  ViewerContext ViewerContext `json:"viewerContext"`
}

type Author struct {
  Fid int64 `json:"fid"`
  Username string `json:"username"`
  DisplayName string `json:"displayName"`
  Pfp PFP `json:"pfp"`
  Profile Profile `json:"profile"`
  FollowerCount int64 `json:"followerCount"`
  FollowingCount int64 `json:"followingCount"`
}

type PFP struct {
  URL string `json:"url"`
  Verified bool `json:"verified"`
}

type Profile struct {
  Bio Bio `json:"bio"`
}

type Bio struct {
  Text string `json:"text"`
}

type Replies struct {
  Count int32 `json:"count"`
}

type Reactions struct {
  Count int32 `json:"count"`
}

type Recasts struct {
  Count int32 `json:"count"`
  Recasters []Recasters `json:"recastes"`
}

type Recasters struct {
  Fid int64 `json:"fid"`
  Username string `json:"username"`
  DisplayName string `json:"displayName"`
  RecastHash string `json:"recastHash"`
}

type ViewerContext struct {
  Reacted bool `json:"reacted"`
  Recast bool `json:"recast"`
  Watched bool `json:"watched"`
}

type Watches struct {
  Count int32 `json:"count"`
}

type Next struct {
  Cursor string `json:"cursor"`
}
