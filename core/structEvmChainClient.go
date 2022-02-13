package core

type EvmChainClient struct {
	ChainName              string   `json:"chainName"`
	ChainID                uint64   `json:"chainID"`
	NativeAsset            string   `json:"nativeAsset"`
	MainRPCAddress         string   `json:"mainRPCAddress"`
	BroadcastRPCAddresses  []string `json:"broadcastRPCAddresses"`
	TradingContractAddress string   `json:"tradingContractAddress"`
	TradingContractVersion string   `json:"tradingContractVersion"`
	RegisteredWorkers      uint32   `json:"registeredWorkers"`
	GasMode                string   `json:"gasMode"`
	GasType                string   `json:"gasType"`
	GasSpecs               []string `json:"gasSpecs"`
	SyncMode               string   `json:"syncMode"`
	BlocksBeforeResync     uint64   `json:"blocksBeforeResync"`
	CallStatic             bool     `json:"callStatic"`
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
