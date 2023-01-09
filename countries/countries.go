package countries

type Country struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IsoAlpha2    string `json:"iso_alpha2"`
	IsoAlpha3    string `json:"iso_alpha3"`
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"curreny_name"`
	CurrencySign string `json:"currency_sign"`
	PhoneCode    string `json:"phone_code"`
}
