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
}
