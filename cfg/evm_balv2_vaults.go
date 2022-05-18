package cfg

import (
	"github.com/codehard-labs/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	EVMBalV2VaultsCollectionName = "evm_balv2_vaults"
)

type BalV2Vault struct {
	Name                string   `bson:"name" json:"name"`
	Source              string   `bson:"source" json:"source"`
	VaultAddress        string   `bson:"vaultAddress" json:"vaultAddress"`
	WeightedPoolFactory string   `bson:"weightedPoolFactory" json:"weightedPoolFactory"`
	WeightedPools       []string `bson:"weightedPools" json:"weightedPools"`
	StablePoolFactory   string   `bson:"stablePoolFactory" json:"stablePoolFactory"`
	StablePools         []string `bson:"stablePools" json:"stablePools"`
	Trading             bool     `bson:"trading" json:"trading"`
}

func GetAllBalV2VaultsOfOneSource(source string) ([]BalV2Vault, error) {
	var res []BalV2Vault
	filter := &bson.M{"source": source}
	err := mongo.FindAll(EVMBalV2VaultsCollectionName, filter, &res)
	return res, err
}
