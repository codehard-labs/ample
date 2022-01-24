package core

type ObexTradingPair struct {
	PairName     string `json:"pairName"`
	PairType     string `json:"pairType"`
	ExchangeName string `json:"exchangeName"`
	QuoteAsset   string `json:"quoteAsset"`
	BaseAsset    string `json:"baseAsset"`
	StepSize     string `json:"stepSize"`
	TakerFee     string `json:"takerFee"`
	Trading      bool   `json:"trading"`
}

func LoadObexTradingPairsFromDB() ([]ObexTradingPair, error) {
	name := "obexTradingPairs"
	data := []ObexTradingPair{}
	//
	err := loadFromDB(name, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveObexTradingPairsToDB(data []ObexTradingPair) error {
	name := "obexTradingPairs"
	//
	return saveToDB(name, &data)
}
