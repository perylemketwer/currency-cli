package models

type Price struct {
	Bid string `json:"bid"`
}

type CoinToReal struct {
	EuroToReal          Price `json:"EURBRL"`
	DolarToReal         Price `json:"USDBRL"`
	PoundToReal         Price `json:"GBPBRL"`
	BitcoinToReal       Price `json:"BTCBRL"`
	DogecoinToReal      Price `json:"DOGEBRL"`
	CanadianDolarToReal Price `json:"CADBRL"`
}
