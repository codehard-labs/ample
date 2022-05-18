package cfg

import (
	"errors"

	"github.com/codehard-labs/ample/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	EVMWorkersCollectionName = "evm_workers"
)

type EVMWorker struct {
	Chain   string `bson:"chain" json:"chain"`
	Index   int    `bson:"index" json:"index"`
	Address string `bson:"address" json:"address"`
	EncPkey string `bson:"encPkey" json:"encPkey"`
	Trading bool   `bson:"trading" json:"trading"`
}

func GetTradingWorkersOfOneEVMChain(chain string, count int) ([]EVMWorker, error) {
	var res []EVMWorker
	filter := &bson.M{"chain": chain, "trading": true}
	ops := options.Find().SetSort(bson.D{{Key: "index", Value: -1}})
	err := mongo.FindAll(EVMWorkersCollectionName, filter, &res, ops)
	if err != nil {
		return nil, err
	}
	if len(res) != count {
		// workers will still be returned even count does not match.
		return res, errors.New("expected workers count does not match")
	}
	return res, nil
}

func GetLastWorkerIndexOfOneEVMChain(chain string) (int, error) {
	var res EVMWorker
	filter := &bson.M{"chain": chain}
	ops := options.FindOne().SetSort(bson.D{{Key: "index", Value: -1}})
	err := mongo.FindOne(EVMWorkersCollectionName, filter, &res, ops)
	if err != nil {
		return 0, err
	}
	return res.Index, err
}
