package core

type EvmChainClient struct {
	ChainName              string   `json:"chainName"`
	ChainID                uint32   `json:"chainID"`
	NativeAsset            string   `json:"nativeAsset"`
	QueryRPCAddress        string   `json:"queryRPCAddress"`
	ExecRPCAddress         string   `json:"execRPCAddress"`
	TradingContractAddress string   `json:"tradingContractAddress"`
	TradingContractVersion string   `json:"tradingContractVersion"`
	RegisteredWorkers      uint32   `json:"registeredWorkers"`
	GasMode                string   `json:"gasMode"`
	GasType                string   `json:"gasType"`
	GasSpecs               []string `json:"gasSpecs"`
}

func LoadEvmChainClientsFromDB() ([]EvmChainClient, error) {
	name := "evmChainClients"
	data := []EvmChainClient{}
	//
	err := loadFromDB(name, &data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveEvmChainClientsToDB(data []EvmChainClient) error {
	name := "evmChainClients"
	//
	return saveToDB(name, &data)
}
