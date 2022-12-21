package user

type CustodyAddressRoot struct {
  Result CustodyAddressResult `json:"result"`
}

type CustodyAddressResult struct {
    CustodyAddress string `json:"custodyAddress"`
}
