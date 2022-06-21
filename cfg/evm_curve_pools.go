package cfg

import (
	"github.com/codehard-labs/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	EVMCurvePoolsCollectionName = "evm_curve_pools"
)

type CurvePool struct {
	Name            string      `bson:"name" json:"name"`
	Source          string      `bson:"source" json:"source"`
	Address         string      `bson:"address" json:"address"`
	LpToken         string      `bson:"lpToken" json:"lpToken"`
	Type            string      `bson:"type" json:"type"`
	Coins           []CurveCoin `bson:"coins" json:"coins"`
	UnderlyingCoins []CurveCoin `bson:"underlyingCoins" json:"underlyingCoins"`
	BasePool        string      `bson:"basePool" json:"basePool"` // Only applies to meta pools
	DynamicFee      bool        `bson:"dynamicFee" json:"dynamicFee"`
	Trading         bool        `bson:"trading" json:"trading"`
}

type CurveCoin struct {
	Address  string `bson:"address" json:"address"`
	Decimals int64  `bson:"decimals" json:"decimals"`
}

func GetAllCurvePoolsOfOneSource(source string) ([]CurvePool, error) {
	var res []CurvePool
	filter := &bson.M{"source": source}
	err := mongo.FindAll(EVMCurvePoolsCollectionName, filter, &res)
	return res, err
}
