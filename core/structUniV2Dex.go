package core

type UniV2Dex struct {
	Name           string `json:"name"`
	Source         string `json:"source"`
	RouterAddress  string `json:"routerAddress"`
	FactoryAddress string `json:"factoryAddress"`
	SubType        string `json:"subType"`
	FeeN           uint64 `json:"feeN"`
	FeeD           uint64 `json:"feeD"`
	Trading        bool   `json:"trading"`
}

func LoadUniV2DexsFromDB() ([]UniV2Dex, error) {
	name := "uniV2Dexs"
	data := []UniV2Dex{}
	//
	err := loadFromDB(name, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveUniV2DexsToDB(data []UniV2Dex) error {
	name := "uniV2Dexs"
	//
	return saveToDB(name, &data)
}
