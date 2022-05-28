package cfg

import (
	"github.com/codehard-labs/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	EVMUniV3DexsCollectionName = "evm_univ3_dexs"
)

type UniV3Dex struct {
	Name           string   `bson:"name" json:"name"`
	Source         string   `bson:"source" json:"source"`
	RouterAddress  string   `bson:"routerAddress" json:"routerAddress"`
	FactoryAddress string   `bson:"factoryAddress" json:"factoryAddress"`
	Pools          []string `bson:"pools" json:"pools"`
	Trading        bool     `bson:"trading" json:"trading"`
}

func GetAllUniV3DexsOfOneSource(source string) ([]UniV3Dex, error) {
	var res []UniV3Dex
	filter := &bson.M{"source": source}
	err := mongo.FindAll(EVMUniV3DexsCollectionName, filter, &res)
	return res, err
}
