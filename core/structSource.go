package core

type Source struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	On      bool   `json:"on"`
	Trading bool   `json:"trading"`
}

func LoadSourcesFromDB() ([]Source, error) {
	name := "sources"
	data := []Source{}
	//
	err := loadFromDB(name, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveSourcesToDB(data []Source) error {
	name := "sources"
	//
	return saveToDB(name, &data)
}
