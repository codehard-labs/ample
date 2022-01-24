package core

// Asset on a specific Source to be considered
type Token struct {
	AssetName  string `json:"assetName"`
	GroupName  string `json:"groupName"`
	Source     string `json:"source"`
	Identifier string `json:"identifier"` // source specific name
	Decimals   uint32 `json:"decimals"`   // only applicable if source is some blockchain
	Trading    bool   `json:"trading"`
}

func LoadTokensFromDB() ([]Token, error) {
	name := "tokens"
	data := []Token{}
	//
	err := loadFromDB(name, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveTokensToDB(data []Token) error {
	name := "tokens"
	//
	return saveToDB(name, &data)
}
