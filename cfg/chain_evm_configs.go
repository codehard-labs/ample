package cfg

import (
	"github.com/codehard-labs/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	EVMChainParamsCollectionName = "chain_evm_configs"
)

type ChainEVMConfig struct {
	ChainName                 string   `bson:"chainName" json:"chainName"`
	ChainID                   uint64   `bson:"chainID" json:"chainID"`
	NativeAsset               string   `bson:"nativeAsset" json:"nativeAsset"`
	NativeAssetDecimals       uint64   `bson:"nativeAssetDecimals" json:"nativeAssetDecimals"`
	MainNodeAddress           string   `bson:"mainNodeAddress" json:"mainNodeAddress"`
	BroadcastAddresses        []string `bson:"broadcastAddresses" json:"broadcastAddresses"`
	TradingContractAddress    string   `bson:"tradingContractAddress" json:"tradingContractAddress"`
	TradingContractVersion    string   `bson:"tradingContractVersion" json:"tradingContractVersion"`
	WorkersCount              int      `bson:"workersCount" json:"workersCount"`
	SyncMode                  string   `bson:"syncMode" json:"syncMode"`
	FilterQuerySupport        string   `bson:"filterQuerySupport" json:"filterQuerySupport"`
	ResyncPeriod              uint64   `bson:"resyncPeriod" json:"resyncPeriod"`
	EstGasInsteadOfCallStatic bool     `bson:"estGasInsteadOfCallStatic" json:"estGasInsteadOfCallStatic"`
	SnipeModeEnabled          bool     `bson:"snipeModeEnabled" json:"snipeMode"`
}

func GetAllChainEVMConfig() ([]ChainEVMConfig, error) {
	var res []ChainEVMConfig
	err := mongo.FindAll(EVMChainParamsCollectionName, mongo.EmptyFilter, &res)
	return res, err
}

func GetChainEVMConfig(chainName string) (ChainEVMConfig, error) {
	var res ChainEVMConfig
	filter := &bson.M{"chainName": chainName}
	err := mongo.FindOne(EVMChainParamsCollectionName, filter, &res)
	return res, err
}
