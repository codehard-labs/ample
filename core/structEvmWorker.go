package core

// Let's not implement this until making connection secured

type EvmWorker struct {
	NickName      string `json:"nickName"`
	Address       string `json:"address"`
	PkeyEncrypted string `json:"pkeyEncrypted"`
}

func LoadEvmWorkerFromDB() ([]EvmWorker, error) {
	name := "evmWorkers"
	data := []EvmWorker{}
	//
	err := loadFromDB(name, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveEvmWorkerToDB(data []EvmWorker) error {
	// this should generally not be allowed
	name := "evmWorkers"
	//
	return saveToDB(name, &data)
}
