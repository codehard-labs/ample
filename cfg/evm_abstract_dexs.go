package cfg

import (
	"github.com/codehard-labs/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	EVMAbstractDexsCollectionName = "evm_abstract_dexs"
)

type EVMAbstractDex struct {
	Name                 string   `bson:"name" json:"name"`
	Source               string   `bson:"source" json:"source"`
	Type                 string   `bson:"type" json:"type"`
	RouterAddress        string   `bson:"routerAddress" json:"routerAddress"`
	FactoryAddress       string   `bson:"factoryAddress" json:"factoryAddress"`
	FactoryDeployedBlock uint64   `bson:"factoryDeployedBlock" json:"factoryDeployedBlock"`
	Pools                []string `bson:"pools" json:"pools"`
	Trading              bool     `bson:"trading" json:"trading"`
}

func GetAllEVMAbstractDexsOfOneSource(source string) ([]EVMAbstractDex, error) {
	var res []EVMAbstractDex
	filter := &bson.M{"source": source}
	err := mongo.FindAll(EVMAbstractDexsCollectionName, filter, &res)
	return res, err
}
