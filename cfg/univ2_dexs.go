package cfg

import (
	"github.com/codehard-labs/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	UniV2DexsCollectionName = "univ2_dexs"
)

type UniV2Dex struct {
	Name           string `bson:"name" json:"name"`
	Source         string `bson:"source" json:"source"`
	RouterAddress  string `bson:"routerAddress" json:"routerAddress"`
	FactoryAddress string `bson:"factoryAddress" json:"factoryAddress"`
	SubType        string `bson:"subType" json:"subType"`
	FeeN           uint64 `bson:"feeN" json:"feeN"`
	FeeD           uint64 `bson:"feeD" json:"feeD"`
	Trading        bool   `bson:"trading" json:"trading"`
}

func GetAllUniV2DexsOfOneSource(source string) ([]UniV2Dex, error) {
	var res []UniV2Dex
	filter := &bson.M{"source": source}
	err := mongo.FindAll(UniV2DexsCollectionName, filter, &res)
	return res, err
}
